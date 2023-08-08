curl -X POST http://192.168.1.101:8070/auth/visitor/sign-up -d \
'{"name":"michael3", "email":"mi@m.com", "password":"qwerty123"}'

curl -X POST http://192.168.1.101:8070/auth/coach/sign-up -d \
'{"name":"michael2coach", "email":"mi@m.com", "password":"qwerty123"}'

curl -X POST http://192.168.1.101:8070/auth/activity/sign-up -d \
'{"name":"veloerg#12345", "description":"A simple Veloergometer", "club_id": 1}'

curl -X POST http://192.168.1.101:8070/auth/activity/sign-in -d \
'{"name": "veloerg#12345"}'

curl -X POST http://192.168.1.101:8070/auth/visitor/sign-in -d \
'{"login":"m", "password":"m"}'

curl -X POST http://192.168.1.101:8070/auth/coach/sign-in -d \
'{"login":"michael2coach", "password":"qwerty123"}'

curl -X POST http://192.168.1.101:8070/api/club/ -d \
'{"address": "м. Харків, вул. Героїв Праці, 20В", "name": "Геркулес"}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzNTc5MTcsImlhdCI6MTY5MDc2NTkxNywiZW50aXR5X2lkIjoyfQ.BMqmZ2NNi_WVqvEaKVUMX_Xp8yruKlO8Lfb6CnLRtUQ'

curl -X GET http://192.168.1.101:8070/api/club/ \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1NzMsImlhdCI6MTY4NzEwNzU3MywiZW50aXR5X2lkIjoxfQ.PUMshw54azu5ikujxdieLzBcxAOPnlM0KtlAbgMySMs'

curl -X GET http://192.168.1.101:8070/api/club/1 \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1NzMsImlhdCI6MTY4NzEwNzU3MywiZW50aXR5X2lkIjoxfQ.PUMshw54azu5ikujxdieLzBcxAOPnlM0KtlAbgMySMs'

curl -X POST http://192.168.1.101:8070/api/phys_info/ -d \
'{"height": 187.56, "weight":67.53}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1MzQsImlhdCI6MTY4NzEwNzUzNCwiZW50aXR5X2lkIjoxfQ.M59wtLTpSpTjJ7ePWpI5A22f916vJ8SHzUfVwdT0efw'

curl -X GET http://192.168.1.101:8070/api/phys_info/ \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM2OTA1NDAsImlhdCI6MTY5MTA5ODU0MCwiZW50aXR5X2lkIjoxM30.LxA1ogfCRUJ1BzwxY_kuGA0W-_Ss_mfkwH55CJ0AyA4'

curl -X PUT http://192.168.1.101:8070/api/phys_info/ \
-d '{"height": "191.25", "weight": "75.0"}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM2OTA1NDAsImlhdCI6MTY5MTA5ODU0MCwiZW50aXR5X2lkIjoxM30.LxA1ogfCRUJ1BzwxY_kuGA0W-_Ss_mfkwH55CJ0AyA4'

curl -X POST http://192.168.1.101:8070/api/training/ \
-d '{"start": "2023-06-09T00:26:35", "end": "2023-06-09T01:30:00", "club_id": 1}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzNTc5MTcsImlhdCI6MTY5MDc2NTkxNywiZW50aXR5X2lkIjoyfQ.BMqmZ2NNi_WVqvEaKVUMX_Xp8yruKlO8Lfb6CnLRtUQ'

curl -X GET http://192.168.1.101:8070/api/training/1 \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzNTc5MTcsImlhdCI6MTY5MDc2NTkxNywiZW50aXR5X2lkIjoyfQ.BMqmZ2NNi_WVqvEaKVUMX_Xp8yruKlO8Lfb6CnLRtUQ'

curl -X POST http://192.168.1.101:8070/api/states_types/ \
-d '{"name": "beats per minute", "description": "measure of heart beats count per minute", "unit": "bpm"}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1NzMsImlhdCI6MTY4NzEwNzU3MywiZW50aXR5X2lkIjoxfQ.PUMshw54azu5ikujxdieLzBcxAOPnlM0KtlAbgMySMs'

curl -X GET http://192.168.1.101:8070/api/states_types/1 \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1NzMsImlhdCI6MTY4NzEwNzU3MywiZW50aXR5X2lkIjoxfQ.PUMshw54azu5ikujxdieLzBcxAOPnlM0KtlAbgMySMs'

curl -X PUT http://192.168.1.101:8070/api/states_types/1 \
-d '{"name": "сердечний ритм", "description": "кількість ударів серця за хвилину", "unit": "раз/хв"}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1NzMsImlhdCI6MTY4NzEwNzU3MywiZW50aXR5X2lkIjoxfQ.PUMshw54azu5ikujxdieLzBcxAOPnlM0KtlAbgMySMs'

curl -X POST http://192.168.1.101:8070/api/activity/ \
-d '{"name": "велоергометр", "description": "-", "club_id": 1}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1NzMsImlhdCI6MTY4NzEwNzU3MywiZW50aXR5X2lkIjoxfQ.PUMshw54azu5ikujxdieLzBcxAOPnlM0KtlAbgMySMs'

curl -X GET http://192.168.1.101:8070/api/activity/3 \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1NzMsImlhdCI6MTY4NzEwNzU3MywiZW50aXR5X2lkIjoxfQ.PUMshw54azu5ikujxdieLzBcxAOPnlM0KtlAbgMySMs'

curl -X DELETE http://192.168.1.101:8070/api/activity/1 \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1NzMsImlhdCI6MTY4NzEwNzU3MywiZW50aXR5X2lkIjoxfQ.PUMshw54azu5ikujxdieLzBcxAOPnlM0KtlAbgMySMs'

curl -X POST http://192.168.1.101:8070/api/attendance/ \
-d '{"training_id": 1}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1MzQsImlhdCI6MTY4NzEwNzUzNCwiZW50aXR5X2lkIjoxfQ.M59wtLTpSpTjJ7ePWpI5A22f916vJ8SHzUfVwdT0efw'

curl -X GET http://192.168.1.101:8070/api/attendance/1 \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1MzQsImlhdCI6MTY4NzEwNzUzNCwiZW50aXR5X2lkIjoxfQ.M59wtLTpSpTjJ7ePWpI5A22f916vJ8SHzUfVwdT0efw'

# curl -X POST http://192.168.1.101:8070/api/activity_usage/ \
# -d '{"activity_id": 2, "start": "2023-06-09T03:22:15", "end": "2023-06-09T03:24:15", "training_id": 1}' \
# -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk2OTk1MzQsImlhdCI6MTY4NzEwNzUzNCwiZW50aXR5X2lkIjoxfQ.M59wtLTpSpTjJ7ePWpI5A22f916vJ8SHzUfVwdT0efw'

curl -X POST http://192.168.1.101:8070/api/activity_usage/ \
-d '{"activity_name": "veloerg#12345", "training_id": 1}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk3NzYzODAsImlhdCI6MTY4NzE4NDM4MCwiZW50aXR5X2lkIjoyfQ.u0tNjgAjiX26b9k6LOqZGTg-6MA6P7t1XaEdvVzW8BY'

curl -X GET http://192.168.1.101:8070/api/activity_usage/?name=veloerg%2312345 \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk3NzYzODAsImlhdCI6MTY4NzE4NDM4MCwiZW50aXR5X2lkIjoyfQ.u0tNjgAjiX26b9k6LOqZGTg-6MA6P7t1XaEdvVzW8BY'

curl -X POST http://192.168.1.101:8070/api/physical_state/ \
-d '{"activity_usage_id": 2, "unit_amount": 153, "state_type_id": 1, "at": "2023-06-09T03:23:16", "secs": 60}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk3NzYzODAsImlhdCI6MTY4NzE4NDM4MCwiZW50aXR5X2lkIjoyfQ.u0tNjgAjiX26b9k6LOqZGTg-6MA6P7t1XaEdvVzW8BY'

curl -X GET http://192.168.1.101:8070/api/physical_state/1 \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk3NzYzODAsImlhdCI6MTY4NzE4NDM4MCwiZW50aXR5X2lkIjoyfQ.u0tNjgAjiX26b9k6LOqZGTg-6MA6P7t1XaEdvVzW8BY'

curl -X POST http://192.168.1.101:8070/api/activity_state/ \
-d '{"unit_amount": 2.69, "state_type_id": 1, "secs": 5}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbnRpdHlOYW1lIjoidmVsb2VyZyMxMjM0NSIsImV4cCI6IjIwMjMtMDctMjBUMDM6MzQ6NTAuNjk3MDA1OCswMzowMCIsImlhdCI6IjIwMjMtMDYtMjBUMDM6MzQ6NTAuNjk3MDA1OCswMzowMCJ9.FELxnXWoPa-IvhwH_uPqfGoMBt2syBCCOQMZkpNSG0M'

curl -X GET http://192.168.1.101:8070/api/activity_state/4 \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbnRpdHlOYW1lIjoidmVsb2VyZyMxMjM0NSIsImV4cCI6IjIwMjMtMDctMjBUMDM6MzQ6NTAuNjk3MDA1OCswMzowMCIsImlhdCI6IjIwMjMtMDYtMjBUMDM6MzQ6NTAuNjk3MDA1OCswMzowMCJ9.FELxnXWoPa-IvhwH_uPqfGoMBt2syBCCOQMZkpNSG0M'

curl -X GET http://192.168.1.101:8070/api/stats/activity_usages?activityName=veloerg%2312345679 \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM5OTE3NjIsImlhdCI6MTY5MTM5OTc2MiwiZW50aXR5X2lkIjozfQ._W6hG8w31qdRkO_7jsBkbYqGccUZC3WkspL1s1zkggg'
