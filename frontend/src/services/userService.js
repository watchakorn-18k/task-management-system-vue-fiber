import { config } from '@/config';

const urlUsers = config.apiUrl + '/api/user';


const authLogin = async (data) => {
    try {
        const response = await fetch(`${urlUsers}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });
        const result = await response.json();
        return result.data;
    } catch (error) {
        console.error('Error:', error);
        throw error;
    }
}

export { authLogin };