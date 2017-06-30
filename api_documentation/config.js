// this is passed to json-schema-example-loader
export default {
  title: 'Notifications Service',
  // enhance your CURL examples by base url and request headers...
  curl: {
    baseUrl: 'https://localhost:3000/v1',
    requestHeaders: {
      required: [
        'Content-Type',
      ],
      properties: {
        'Content-Type': {
          type: 'string',
          enum: [
            'application/json',
          ],
          example: 'application/json',
          description: 'Content type of the API request',
        },
      },
    },
  },
};
