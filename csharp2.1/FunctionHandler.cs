using System;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using System.Net;
using System.IO;
using System.Collections.Generic;
using Newtonsoft.Json;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;

namespace Function
{
    public class FunctionHandler
    {
        public async Task Handle(HttpContext context)
        {
            // Get the files array
            var files = context.Request.Form.Files;
            // Create a path to save the file
            var filePath = "/tmp/";
            var fileContentType = (string)null;

            foreach (var file in files)
            {
                // Only get the image file
                if (file.Length > 0 && file.Name == "image")
                {
                    // Get the contentType
                    fileContentType = file.ContentType;
                    // Add the filename to the path
                    filePath += file.FileName;
                    // Save the file to disk
                    using (var fileStream = new FileStream(filePath, FileMode.Create))
                    {
                        await file.CopyToAsync(fileStream);
                    }       
                }
            }

            // Send the file as a response
            context.Response.Clear();
            context.Response.StatusCode = 200;
            context.Response.ContentType = fileContentType;
            await context.Response.SendFileAsync(filePath);

            return;
        }
    }
}
