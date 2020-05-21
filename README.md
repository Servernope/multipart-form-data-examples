<img  src="http://d1bxssiwxq2vbt.cloudfront.net/logo-full.png"  alt="Servernope logo"  title="Servernope"  align="right"  height="72"  width="240"/>

## Handling multipart form-data requests

This template receives an image via a `multipart form-data` request, saves it to disk, and then uses the image as a response.

To test this code, [create a hosted function](https://www.servernope.com/docs/Selling-Cloud-Functions/Functions/Creating-a-function), [configure the routes](https://www.servernope.com/docs/Selling-Cloud-Functions/Functions/Function-routes) and [parameters](https://www.servernope.com/docs/Selling-Cloud-Functions/Functions/Function-parameters) as shown below, [clone the code to your machine](https://www.servernope.com/docs/Selling-Cloud-Functions/CLI-Tool/Pull-your-function-code), replace the contents of your `/function` folder by its corespondent runtime and [test your function locally](https://www.servernope.com/docs/Selling-Cloud-Functions/CLI-Tool/Testing-locally).

#### Route:
![Route](https://servernope-attachments.s3.amazonaws.com/gist/multipart-form-data-example-1.png)
#### Parameters:
![Parameters](https://servernope-attachments.s3.amazonaws.com/gist/multipart-form-data-example-2.png)
