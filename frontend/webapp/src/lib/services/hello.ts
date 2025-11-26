import { baseURL } from '$lib/config';

export interface HelloResponse {
	time: string;
}

export async function fetchCurrentTime(): Promise<HelloResponse> {
	const response = await fetch(`${baseURL()}/api/hello`);

	if (!response.ok) {
		throw new Error(`Failed to fetch time: ${response.statusText}`);
	}

	return response.json();
}