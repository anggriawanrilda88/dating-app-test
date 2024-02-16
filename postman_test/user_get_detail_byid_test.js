// Set environment variable for the base URL
const baseURL = "http://localhost/dating-app-test/api/v1/users/1";
const token = "your_token";


// Request to get user by ID endpoint
pm.sendRequest({
    url: `${baseURL}`,
    method: 'GET',
    header: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + token // Replace {{your_token}} with the actual Bearer token
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
        pm.response.to.have.jsonBody('message', 'Successfully get user');
        pm.response.to.have.jsonBody('data.id');
        pm.response.to.have.jsonBody('data.email');
        pm.response.to.have.jsonBody('data.status');
        pm.response.to.have.jsonBody('data.createdAt');
        pm.response.to.have.jsonBody('data.updatedAt');
    });
});