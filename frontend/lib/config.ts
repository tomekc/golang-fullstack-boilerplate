
import { dev } from '$app/environment';
import devConfig from './devconfig.json'


export function baseURL() : string {
    if (dev) {
        return `http://localhost:${devConfig.backend_port}`;
    } else {
        return ''
    }
}