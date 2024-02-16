// Set environment variable for the base URL
const baseURL = "http://localhost/dating-app-test/api/v1/users";

// Request to register user endpoint
pm.sendRequest({
    url: `${baseURL}/`,
    method: 'POST',
    header: {
        'Content-Type': 'application/json',
    },
    body: {
        mode: 'raw',
        raw: JSON.stringify({
            "email": "your_username",
            "password": "your_password"
        })
    }
}, (err, response) => {
    // Check if there was no error during the request
    pm.test("No request error", function () {
        pm.expect(err).to.be.null;
    });

    // Check if the response status is 200 OK
    pm.test("Status code is 200", function () {
        pm.response.to.have.status(200);
    });

    // Check if the response is in JSON format
    pm.test("Response is JSON", function () {
        pm.response.to.have.header("Content-Type", "application/json; charset=utf-8");
    });

    // Check the properties in the response body
    pm.test("Response contains expected properties", function () {
        pm.response.to.have.jsonBody('message', 'User successfully registered');
    });
});