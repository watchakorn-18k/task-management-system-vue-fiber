import { config } from '@/config';
import { getCookieUserJWT } from "@/utils/cookies";

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
        if (!response.ok) {
            return result.message
        }
        return result.data;
    } catch (error) {
        console.error('Error:', error);
        throw error;
    }
}

const registerUsers = async (data) => {
    try {
        const response = await fetch(`${urlUsers}/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });
        const result = await response.json();
        if (!response.ok) {
            return result.message
        }
        return result.data;
    } catch (error) {
        console.error('Error:', error);
        throw error;
    }
}

const getProfile = async () => {
    try {
        const response = await fetch(`${urlUsers}/get_profile`, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${getCookieUserJWT()}`,
            },
        });
        const result = await response.json();
        if (!response.ok) {
            return result.message
        }
        return result.data;
    } catch (error) {
        console.error('Error:', error);
        throw error;
    }
}

export { authLogin, registerUsers, getProfile };