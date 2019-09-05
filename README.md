# morningglory

points, and trophy management system.

## api/user

### `POST /api/users/{USERNAME}`

register user.

#### Request body

```js
{
    'token': 'secret kyphrase'
}
```


#### Response

* 400
    * `{ 'message': 'USERNAME was already registered' }`
* 201
    * `{ 'message': 'user \'USERNAME\' was created' }`



### `PUT /api/users/{USERNAME}`

update the user token.

#### Request header

* `X-USER-TOKEN` [required]
    * user token.

#### Request body

```js
{
    'newToken': 'secret new token'
}
```

#### Response

* 400
    * `{ 'message': 'token did not match' }`, 
* 404
    * `{ 'message': 'user \'USERNAME\' did not found' }`
* 200
    * `{ 'message': 'success' }`


### `DELETE /api/users/{USERNAME}`

#### Request header

* `X-USER-TOKEN` [required]
    * user token.

#### Response

* 400
    * `{ 'message': 'token did not match' }`, 
* 404
    * `{ 'message': 'user \'USERNAME\' did not found' }`
* 200
    * `{ 'message': 'success' }`

## api/points

### `POST /api/users/{USERNAME}/points`

#### Request header

* `X-USER-TOKEN` [required]
    * user token.
    
#### Request body

```js
{
  'repository': 'USER/REPO',
  'action': 'commit', // commit, push, pullrequest, issue, comments, others
  'ref_url': 'https://....', // url of action.
}
```

#### Response

* 400
    * `{ 'message': 'token did not match' }`, 
* 404
    * `{ 'message': 'user \'USERNAME\' did not found' }`
* 200
    * `{ 'message': 'success' }`

## api/webhooks

### `GET /api/users/{USERNAME}/webhooks`

### `POST /api/users/{USERNAME}/webhooks`

### `POST /api/users/{USERNAME}/webhooks/{WEBHOOK_HASH}`

invoke the webhook.

### `DELETE /api/users/{USERNAME}/webhooks/{WEBHOOK_HASH}`

