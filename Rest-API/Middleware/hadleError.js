function errorHandler(err, req, res, next) {
    console.log(err)
    let status = 500
    let reason = {
        error: "Internal Server Error",
    }

    if (
        err.name === "SequelizeUniqueConstraintError" ||
        err.name === "SequelizeValidationError" ||
        err.name === "SequelizeDatabaseError"
    ) {
        status = 401
        reason = { name: err.message }
    } else if (
        err.name === "InvalidUser" ||
        err.name === "InvalidCreated" ||
        err.name === "ErrorUpdateCategory" 
    ) {
        status = 400
        reason = { name: err.message }
    }

    res.status(status).json(reason)
}

module.exports = errorHandler
