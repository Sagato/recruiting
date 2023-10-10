package user_storage

var GetUsersStatement = `select * from user_service.users`
var GetUserByIdStatement = `select * from user_service.users where id = %d`
var InsertUserStatement = `
INSERT INTO user_service.users ()
VALUES ()`
var UpdateUserStatement = `UPDATE user_service.users 
SET ... WHERE id = $4;`
var DeleteUserByIdStatement = `delete from user_service.users where id = %d RETURNING id`
