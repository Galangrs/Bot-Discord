if (process.env.NODE_ENV !== 'production') {
    require('dotenv').config();
}

const express = require('express')
const cors = require('cors')
const app = express()
const port = process.env.PORT || 4000

app.use(cors())

app.use(express.urlencoded({ extended: false }));
app.use(express.json())

const router = require('./Routers/router.js')
const errorHandler = require("./Middleware/hadleError.js");

app.use("/",router)

app.use(errorHandler)

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})


