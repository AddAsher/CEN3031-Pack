import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';


export interface User {
    username: string;
    password: string;
}

export interface Club {
    description: string;
    leader: string;
    contact: string;
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

    register(username: string, password: string) {
        const url = `${this.baseUrl}/registration`;
        const body = { username, password };
        return this.http.post<User>(url, body);
    }

    getClubs(): Observable<Map<string, Club>> {
        const url = `${this.baseUrl}/home`;
        return this.http.get<Map<string, Club>>(url);
    }
}