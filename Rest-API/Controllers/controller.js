const { User } = require('../models/index.js');

class Controller {
    static async register(req,res,next){
        const { id:userid } = req.body
        try {
            if (!userid) throw { name: "InvalidCreated", message:"id cannot be null" }
            const response = await User.create({
                userid,
                ballance:0
            });
            if (!response) throw { name: "InvalidCreated", message:"Account already register" }
            res.status(201).json({
                name:"Success register account"
            })
        } catch (error) {
            next(error)
        }
    }
}

module.exports = Controller