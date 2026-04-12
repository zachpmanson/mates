export function params(v: Record<string, string>) {
	return new URLSearchParams(v).toString();
}
