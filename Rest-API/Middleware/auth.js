function Authentication(req,res,next){
    const { id:userid } = req.headers
    try {
        if (!userid) throw { name:"Authentication",message: "Authentication" }
        req.user = {
            id:userid
        }
        next()
    } catch (error) {
        next(error)
    }
}

function Authorization(req,res,next){
    const { id:userid } = req.headers
    try {
        if (userid != process.env.DISCORD_ID_OWNER) throw { name:"Authorization",message: "Authorization" }
        next()
    } catch (error) {
        next(error)
    }
}
module.exports = {
    Authentication,
    Authorization
}