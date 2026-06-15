#!/usr/bin/env python3
"""
adc_cli.py — run a single NetScaler (Citrix ADC) CLI command over SSH and print
clean output. Handles the interactive password prompt and `man`/`--More--`
pagination using only the Python standard library (no sshpass/expect/paramiko).

Usage:
    python3 adc_cli.py "<cli command>"

Examples:
    python3 adc_cli.py "man add nslicenseserver"
    python3 adc_cli.py "add nslicenseserver ?"
    python3 adc_cli.py "show nslicenseserver ?"

Connection settings (all overridable via environment variables):
    NS_SSH_HOST  (default: 10.101.132.151, or derived from NS_URL host)
    NS_SSH_USER  (default: $NS_LOGIN or "nsroot")
    NS_SSH_PASS  (default: $NS_PASSWORD or "CADS123$%^")

The NetScaler CLI accepts the concatenated NITRO resource name as a single
token (e.g. `add nslicenseserver` resolves to `add ns licenseserver`), so you
can pass the NITRO resource name verbatim.
"""
import os
import pty
import re
import select
import sys
import time
from urllib.parse import urlparse


def resolve_conn():
    host = os.environ.get("NS_SSH_HOST")
    if not host:
        ns_url = os.environ.get("NS_URL")
        if ns_url:
            netloc = urlparse(ns_url if "://" in ns_url else "http://" + ns_url).netloc
            host = netloc.split("@")[-1].split(":")[0]
    host = host or "10.101.132.151"
    user = os.environ.get("NS_SSH_USER") or os.environ.get("NS_LOGIN") or "nsroot"
    passwd = os.environ.get("NS_SSH_PASS") or os.environ.get("NS_PASSWORD") or "CADS123$%^"
    return host, user, passwd


ANSI_RE = re.compile(rb"\x1b\[[0-9;?]*[a-zA-Z]|\x1b[=>]|\x1b\][^\x07]*\x07|\r")
MORE_RE = re.compile(rb"--More--\([^)]*\)|No next tag\s+\(press RETURN\)|\(END\)")


def clean(raw: bytes) -> str:
    raw = ANSI_RE.sub(b"", raw)
    raw = MORE_RE.sub(b"", raw)
    text = raw.decode("utf-8", "replace")
    # Collapse runs of blank lines left behind by stripped pager markers.
    lines = [ln.rstrip() for ln in text.splitlines()]
    out, blank = [], False
    for ln in lines:
        if ln.strip() == "":
            if not blank:
                out.append("")
            blank = True
        else:
            out.append(ln)
            blank = False
    return "\n".join(out).strip("\n")


def run(cmd: str, timeout: int = 90) -> str:
    host, user, passwd = resolve_conn()
    argv = [
        "ssh",
        "-o", "StrictHostKeyChecking=no",
        "-o", "UserKnownHostsFile=/dev/null",
        "-o", "PreferredAuthentications=password",
        "-o", "PubkeyAuthentication=no",
        "-o", "ConnectTimeout=20",
        f"{user}@{host}",
    ]

    pid, fd = pty.fork()
    if pid == 0:
        os.environ["TERM"] = "dumb"  # suppress bold/pager escape codes
        os.execvp("ssh", argv)
        os._exit(1)

    captured = b""          # everything after the command echo (the answer)
    win = b""               # sliding window for prompt/pager detection
    sent_pass = sent_cmd = done = False
    start = time.time()

    def w(s: str):
        os.write(fd, s.encode())

    while True:
        if time.time() - start > timeout:
            break
        r, _, _ = select.select([fd], [], [], 0.5)
        if r:
            try:
                data = os.read(fd, 4096)
            except OSError:
                break
            if not data:
                break
            win = (win + data)[-512:]
            if sent_cmd:
                captured += data
            low = win.lower()

            if not sent_pass and b"password:" in low:
                w(passwd + "\n")
                sent_pass = True
                time.sleep(0.8)
                win = b""
                continue
            if sent_pass and not sent_cmd and win.rstrip().endswith(b">"):
                time.sleep(0.3)
                w(cmd + "\n")
                sent_cmd = True
                win = b""
                captured = b""
                continue
            if sent_cmd:
                if b"--more--" in low:
                    w(" ")
                    win = b""
                    continue
                if b"press return" in low or b"(end)" in low:
                    w("\r")
                    win = b""
                    continue
        else:
            if sent_cmd and not done and win.rstrip().endswith(b">"):
                w("exit\n")
                done = True
            elif done:
                break
            elif not sent_pass and (time.time() - start) > 25:
                break

    try:
        os.close(fd)
    except OSError:
        pass

    text = clean(captured)
    # Drop the trailing prompt/exit lines.
    text = re.sub(r"\n?>?\s*exit\s*\nBye!.*$", "", text, flags=re.S)
    text = re.sub(r"\n>\s*$", "", text)
    return text.strip("\n")


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print(__doc__)
        sys.exit(2)
    print(run(sys.argv[1]))
