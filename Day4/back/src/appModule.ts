import { MiddlewareConsumer, Module, NestModule } from '@nestjs/common';
import NoteModule from './notes/note.module';

import AppLogger from './appLogger';
import { CorsMiddleware } from './appCors';

@Module({
	imports: [NoteModule],
})
export default class AppModule implements NestModule {
	configure(consumer: MiddlewareConsumer): void {
		consumer.apply(AppLogger).forRoutes('*');
		consumer.apply(CorsMiddleware).forRoutes('*');
	}
}
