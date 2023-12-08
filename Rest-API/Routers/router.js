const express = require("express")
const Controller = require("../Controllers/controller.js")

const router = express.Router()

router.post("/register",Controller.register)


module.exports = router