from flask import Flask, request, send_from_directory

def handle(app):
    # Save the file path
    file_name = ""

    if request.files:
        for file in request.files:
            # Only get the image file
            if file == "image":
                temp_file = request.files.get(file)
                file_name = temp_file.filename
                # Save the file to disk
                temp_file.save("/tmp/" + file_name)

    # Return the file
    return send_from_directory("/tmp/", file_name)
