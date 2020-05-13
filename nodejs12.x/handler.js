"use strict"

const fs = require('fs').promises

module.exports = async (req, res) => {
    // Filter the files array
    const files = req.files.filter(element => element.fieldname == 'image')

    // Check if the image file exists
    if (!files || files.length == 0) {
        return res.status(400).send({
            message: 'Missing image file.'
        })
    }

    // Get the image file object
    const imageFile = files[0]
    // Create a path for the file
    const filePath = `/tmp/${imageFile.originalname}`

    // Save the file to disk
    try {
        await fs.writeFile(filePath, imageFile.buffer)
    } catch (err) {
        console.error(err)
        return res.status(400).send({
            message: `Failed to save file due to error: ${err.message}.`
        })
    }

    // Return the file from disk
    return res.status(200).sendFile(filePath)
}
