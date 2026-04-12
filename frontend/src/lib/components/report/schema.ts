import { z } from 'zod';

export const formSchema = z.object({
	title: z.string().optional(),
	summary: z.string().min(1, 'Summary is required'),
	lat: z.coerce.number(),
	long: z.coerce.number()
});
export type FormSchema = z.infer<typeof formSchema>;
