resource "citrixadc_nscqaparam" "tf_nscqaparam" {
    harqretxdelay = 5
    net1label = "2g"
    minrttnet1 = 25
    lr1coeflist = "intercept=4.95,thruputavg=5.92,iaiavg=-189.48,rttmin=-15.75,loaddelayavg=0.01,noisedelayavg=-2.59"
    lr1probthresh = 0.2
    net1cclscale = "25,50,75"
    net1csqscale = "25,50,75"
    net1logcoef = " 1.49,3.62,-0.14,1.84,4.83"
    net2label = "3g"
    net3label = "4g"
}