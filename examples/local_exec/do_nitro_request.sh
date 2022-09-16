
echo "Set prompt"
curl \
--progress-bar \
-H "Content-Type: application/json" \
-H "X-NITRO-USER: nsroot" \
-H "X-NITRO-PASS: $NITRO_PASS" \
-X PUT \
-d "{\"systemparameter\": { \"promptstring\": \"$PROMPT\"}}" \
-k \
https://$NSIP/nitro/v1/config/systemparameter


echo "add lbvserver 'my_lbvserver'"
curl \
--progress-bar \
-H "Content-Type: application/json" \
-H "X-NITRO-USER: nsroot" \
-H "X-NITRO-PASS: $NITRO_PASS" \
-X POST \
-d "{\"lbvserver\": { \"name\": \"$LBVSERVER\",\"servicetype\": \"ANY\"}}" \
-k \
https://$NSIP/nitro/v1/config/lbvserver


echo "add service 'my_service'"
curl \
--progress-bar \
-H "Content-Type: application/json" \
-H "X-NITRO-USER: nsroot" \
-H "X-NITRO-PASS: $NITRO_PASS" \
-X POST \
-d "{\"service\": { \"name\": \"$SERVICE\",\"ip\": \"10.2.10.1\",\"servicetype\": \"ANY\",\"port\": \"80\"}}" \
-k \
https://$NSIP/nitro/v1/config/service


echo "bind lbvserver 'my_lbvserver' to service 'my_service'"
curl \
--progress-bar \
-H "Content-Type: application/json" \
-H "X-NITRO-USER: nsroot" \
-H "X-NITRO-PASS: $NITRO_PASS" \
-X PUT \
-d "{\"lbvserver_service_binding\": { \"name\": \"$LBVSERVER\",\"servicename\": \"$SERVICE\"}}" \
-k \
https://$NSIP/nitro/v1/config/lbvserver_service_binding


echo "GET binding details of lbvserver 'my_lbvserver'"
curl \
--progress-bar \
-H "Content-Type: application/json" \
-H "X-NITRO-USER: nsroot" \
-H "X-NITRO-PASS: $NITRO_PASS" \
-X GET \
-k \
https://$NSIP/nitro/v1/config/lbvserver_service_binding/$LBVSERVER
