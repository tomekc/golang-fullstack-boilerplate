
import { dev } from '$app/environment';

export function baseURL() : string {
    if (dev) {
        return 'http://localhost:3001';
    } else {
        return ''
    }
}