import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

export interface User {
    username: string;
    password: string;
}

@Injectable({
    providedIn: 'root'
})
export class AuthService {
    private baseUrl = 'http://localhost:8080';

    constructor(private http: HttpClient) { }

    login(username: string, password: string) {
        const url = `${this.baseUrl}/login`;
        const body = { username, password };
        return this.http.post<User>(url, body);
    }
}