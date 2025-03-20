import { Request, Response, NextFunction } from 'express';
import jwt from 'jsonwebtoken';

export interface AuthenticatedRequest extends Request
{
    user?: any;
}

export const authMiddleware = (req: AuthenticatedRequest, res: Response, next: NextFunction) =>
{
    const authHeader = req.headers.authorization;
    if (!authHeader) {
        return res.status(401).json({ error: 'Missing Authorization header' });
    }

    const parts = authHeader.split(' ');
    if (parts.length !== 2 || parts[0] !== 'Bearer') {
        return res.status(401).json({ error: 'Invalid Authorization header format' });
    }

    const token = parts[1];
    try {
        // Para RS256, en lugar de 'JWT_SECRET' deberás obtener la clave pública
        const decoded = jwt.verify(token, process.env.JWT_SECRET || 'tu_clave_secreta');

        req.user = decoded;

        // Opcional: propagar el sub en un header personalizado para los microservicios
        if (decoded && typeof decoded === 'object' && decoded.sub) {
            req.headers['x-user-sub'] = decoded.sub;
        }

        next();
    } catch (error) {
        return res.status(401).json({ error: 'Invalid token' });
    }
};
