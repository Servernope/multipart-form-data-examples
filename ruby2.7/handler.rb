class Handler
  def run(controller, request)
    # Only get the image file
    tempfile = request.params['image'][:tempfile]
    # Get the original file name
    filename = request.params['image'][:filename]
    # Create a new path
    filepath = "/tmp/#{filename}"
    # Copy the file to our path
    File.open(filepath, 'wb') {|f| f.write tempfile.read }

    # Return the file
    controller.send_file filepath
  end
end
