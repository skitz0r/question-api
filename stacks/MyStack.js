import * as sst from "@serverless-stack/resources";

export default class MyStack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);

    // Create a HTTP API
    const api = new sst.Api(this, "Api", {
      routes: {
        "POST /api/user": "cmd/status",
        "POST /api/question": "cmd/questions",
      }
    });

    // Show the endpoint in the output
    this.addOutputs({
      "ApiEndpoint": api.url,
    });
  }
}
