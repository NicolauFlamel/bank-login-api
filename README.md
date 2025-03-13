# Members

Jéferson Alan Salomon
Carlos Manoel Wendorff

# Bank-Login-Api

This is an API for a bank authentication app.

## Running the app

1. Pull the repo.
2. Run `docker-compose up` to deploy the postgres database.
3. From the root run `go run ./cmd/server`
4. On the first run, at `cmd/server/api.go` set the second parameter of the function `SetupDbUsers(db, true)` to `true`, and on subsequent runs let it false.

# Endpoints

- GET `/create-session`:
  - returns: {session_id: string, layout: string}
    - session_id: the id of the session
    - layout: AES-256 encrypted keyboard layout.   
- POST `validate`:
  - request: {session_id: string, sequence: string}
    - session_id: the id of the session
    - sequence: AES-256 encrypted keys typed by the user
  - response: {user, is_valid
   
## Validations

- Session: The application evaluates if the session_id given is the same when returned by the client. Furthermore, evalueates if it's expired (5 min expiration limit) and if it's active (has not been used before).
- Layout: The layout generated will not repeat itself in a hundred runs, it does so by checking in the sessions table if in the last hundred rows this layout preset exists (in case it does, runs the random algorithm again).
- Sequence: The retuned sequence will first be evaluated if its format is valid and the keys it composes of are present in the session's layout.
- Password: The passwords are 4 digit long and are stored in the users table, each digit being one column containing its hash.
- Password Validation: As there is no user validation, the application iterates over all users and check if any of the hashes of both digits of the pair match the users correspondent digit and returns the user it matched.

## Problems

Due to some sort of node and browser difference, we weren't able to decrypt the AES message on the frontend, therefore we did not implement it.
