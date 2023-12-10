const { User, sequelize, CID } = require('../models/index.js');
const { Op } = require('sequelize');

class Controller {
    static async register(req, res, next) {
        const { id: userid } = req.body;
        try {
            if (!userid) throw { name: "InvalidCreated", message: "id cannot be null" };
            const response = await User.create({
                userid,
                balance: 0,
            });
            if (!response) throw { name: "InvalidCreated", message: "Account already registered" };
            res.status(201).json({
                name: "Success register account",
            });
        } catch (error) {
            next(error);
        }
    }

    static async buyCID(req, res, next) {
        const { id:userid } = req.user
        const { quantity: count } = req.body;
        const t = await sequelize.transaction();
        try {
            if (!count) throw { name: "InvalidBuy", message: "quantity cannot be null" };
            if (!Number(count)) throw { name: "InvalidBuy", message: "quantity is integer" };
            if (count < 1) throw { name: "InvalidBuy", message: "Invalid quantity. Please provide a quantity greater than or equal to 1." };
            if (!userid) throw { name: "InvalidBuy", message: "id cannot be null" };

            const { balance } = await User.findOne({
                where: {
                    userid,
                },
            });
            if (balance < count) throw { name: "InvalidBuy", message: "balance is low" };

            await User.update(
                {
                    balance: Math.floor(balance - count),
                },
                {
                    where: { userid },
                    transaction: t,
                },
            );

            const response = await CID.findAll({
                limit: count,
                order: [
                    ['id', 'ASC'],
                ],
                transaction: t,
            });

            // Check if items are sold-out
            if (response.length === 0) {
                throw { name: "InvalidBuy", message: "items are sold-out" };
            } else if (response.length < count){
                throw { name: "InvalidBuy", message: `stock availabel in ${response.length}` };
            }

            await CID.destroy({
                where: {
                    id: {
                        [Op.in]: response.map(cid => cid.id),
                    },
                },
                transaction: t,
            });            

            await t.commit();
            res.status(200).json({ name: response.map(cid => cid.uniqcode).join("\n") });
        } catch (error) {
            await t.rollback();
            next(error);
        }
    }

    static async addBal(req,res,next){
        const { value: count ,id:userid} = req.body;
        try {
            if (!count) throw { name: "InvalidAddBal", message: "value cannot be null" };
            if (!Number(count)) throw { name: "InvalidAddBal", message: "value is integer" };
            if (count < 1) throw { name: "InvalidAddBal", message: "Invalid value. Please provide a value greater than or equal to 1." };
            if (!userid) throw { name: "InvalidAddBal", message: "id cannot be null" };
            
            const response = await User.findOne({
                where: {
                    userid,
                },
            });
            if (!response) throw {name:"InvalidAddBal",message:"account not register"}
            await User.update(
                {
                    balance: Math.floor(response.balance + count),
                },
                {
                    where: { userid },
                },
            );
            res.status(200).json({name:"Success add balance"})
        } catch (error) {
            next(error)
        }
    }

    static async delBal(req,res,next){
        const { value: count ,id:userid} = req.body;
        try {
            if (!count) throw { name: "InvalidDelBal", message: "value cannot be null" };
            if (!Number(count)) throw { name: "InvalidDelBal", message: "value is integer" };
            if (count < 1) throw { name: "InvalidDelBal", message: "Invalid value. Please provide a value greater than or equal to 1." };
            if (!userid) throw { name: "InvalidDelBal", message: "id cannot be null" };
            
            const response = await User.findOne({
                where: {
                    userid,
                },
            });
            if (!response) throw {name:"InvalidAddBal",message:"account not register"}

            if (response.balance < count) {
                throw { name: "InvalidDelBal", message: `balance now ${balance} you cannot deduct in ${count}` };
            }

            await User.update(
                {
                    balance: Math.floor(response.balance - count),
                },
                {
                    where: { userid },
                },
            );
            res.status(200).json({name:"Success deduct balance"})
        } catch (error) {
            next(error)
        }
    }

    static async getBal(req,res,next){
        const { id:userid } = req.user
        try {
            const response = await User.findOne({
                where:{
                    userid
                }
            })
            if (!response) throw {name:"InvalidGetBal",message:"account not register"}

            res.status(200).json({name:"you balace : " + response.balance})
        } catch (error) {
            next(error)
        }
    }
}

module.exports = Controller;
