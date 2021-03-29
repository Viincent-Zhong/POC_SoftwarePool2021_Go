import { NestMiddleware, Injectable } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';

@Injectable()
export class CorsMiddleware implements NestMiddleware {
	use(req: Request, res: Response, next: NextFunction): void {
		res.header('Access-Control-Allow-Origin', '*');
		res.header(
			'Access-Control-Allow-Methods',
			'GET, POST, OPTIONS, PUT, PATCH, DELETE',
		);
		res.header(
			'Access-Control-Allow-Headers',
			'x-access-token, Origin, X-Requested-With, Content-Type, Accept',
		);

		next();
	}
}
