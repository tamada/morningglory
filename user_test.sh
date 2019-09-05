echo -n 'register user "tamada": '
curl -X POST -d '{ token: "hogehoge" }' https://morningglory.appspot.com/v1/users/tamada

echo -n 'update key phrase, but user token not match: '
curl -X PUT -H 'X-USER-TOKEN: NotMatch' -d '{ token: "fugafuga" }' https://morningglory.appspot.com/v1/users/tamada
echo -n 'update key phrase, but user not found: '
curl -X PUT -H 'X-USER-TOKEN: unknwon'  -d '{ token: "fugafuga" }' https://morningglory.appspot.com/v1/users/unknown
echo -n 'update key phrase: '
curl -X PUT -H 'X-USER-TOKEN: hogehoge' -d '{ token: "fugafuga" }' https://morningglory.appspot.com/v1/users/tamada

echo -n 'delete user "tamada", but token not match: '
curl -X DELETE -H 'X-USER-TOKEN: hogehoge' https://morningglory.appspot.com/v1/users/tamada
echo -n 'delete user "tamada", but user not found: '
curl -X DELETE -H 'X-USER-TOKEN: NotMatch' https://morningglory.appspot.com/v1/users/unknwon
echo -n 'delete user "tamada": '
curl -X DELETE -H 'X-USER-TOKEN: fugafuga' https://morningglory.appspot.com/v1/users/tamada
