const express = require('express');
const axios = require('axios');
require('dotenv').config();
const cors = require('cors');

const app = express();
app.use(express.json());
app.use(cors());

// Endpoint to generate an image
app.post('/api/generate-image', async (req, res) => {
    try {
        const response = await axios.post('https://api.openai.com/v1/images/generations', {
            prompt: req.body.prompt,
            n: 1,
            size: "1024x1024",
            quality: "standard",
        }, {
            headers: {
                'Authorization': `Bearer ${process.env.OPENAI_API_KEY}`
            }
        });

        const imageUrl = response.data.data[0].url;
        res.json({ imageUrl });
    } catch (error) {
        console.error('Error generating image:', error);
        res.status(500).json({ message: 'Error generating image' });
    }
});

const PORT = process.env.PORT || 3001;
app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
});
