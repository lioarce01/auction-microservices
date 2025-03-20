import { expressjwt } from "express-jwt";
import jwks from "jwks-rsa";
import { Request, Response, NextFunction } from "express";
import { config } from "dotenv"

config()

console.log("AUTH0_DOMAIN:", process.env.AUTH0_DOMAIN);
console.log("AUTH0_AUDIENCE:", process.env.AUTH0_AUDIENCE);

export const authMiddleware = expressjwt({
    secret: jwks.expressJwtSecret({
        cache: true,
        rateLimit: true,
        jwksRequestsPerMinute: 5,
        jwksUri: `https://${process.env.AUTH0_DOMAIN}/.well-known/jwks.json`,
    }) as any,
    audience: process.env.AUTH0_AUDIENCE,
    issuer: `https://${process.env.AUTH0_DOMAIN}/`,
    algorithms: ["RS256"],
});

export const addUserSubMiddleware = (req: Request, res: Response, next: NextFunction) =>
{
    if (!req.user || !req.user.sub) {
        return res.status(401).json({ error: "User ID not found in token" });
    }

    req.headers["x-user-sub"] = req.user.sub;
    next();
};
