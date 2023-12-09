function errorHandler(err, req, res, next) {
    console.log(err)
    let status = 500
    let reason = {
        name: "Internal Server Error",
    }

    if (
        err.name === "SequelizeUniqueConstraintError" ||
        err.name === "SequelizeValidationError"
    ) {
        status = 401
        reason = { name: err.message }
    } else if (
        err.name === "InvalidBuy" || 
        err.name === "InvalidAddBal" ||
        err.name === "InvalidDelBal"
    ) {
        status = 400
        reason = { name: err.message }
    } else if (
        err.name === "InvalidCreated"
    ) {
        status = 401
        reason = { name: err.message }
    } else if (
        err.name === "SequelizeDatabaseError"
    ) {
        status = 500
        reason = { name: err.message }
    }

    res.status(status).json(reason)
}

module.exports = errorHandler
