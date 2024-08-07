import { writable } from 'svelte/store';
import type { MateReport } from '../types/report';

function createReportStore() {
	const { subscribe, set, update } = writable<{
		reports: MateReport[];
		highlightedReport: number | null;
	}>({
		reports: [],
		highlightedReport: null
	});

	return {
		subscribe,
		set,
		update,
		setHighlightedReport: (report: number | null) =>
			update((o) => ({ ...o, highlightedReport: report })),
		setReports: (reports: MateReport[]) => update((o) => ({ ...o, reports: reports }))
	};
}

export const reportStore = createReportStore();
