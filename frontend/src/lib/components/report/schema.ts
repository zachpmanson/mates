import { z } from 'zod';

export const formSchema = z.object({
	description: z.string().min(1),
	name: z.string(),
	transport: z.enum(['tram', 'train', 'bus'])
});
export type FormSchema = z.infer<typeof formSchema>;
