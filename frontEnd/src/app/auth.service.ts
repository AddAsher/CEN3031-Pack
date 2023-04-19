import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { AnyCatcher } from 'rxjs/internal/AnyCatcher';


export interface User {
    username: string;
    password: string;
}

export interface Club {
    description: string;
    leader: string;
    contact: string;
    hyperlink: string
}

export type NullableClub = Club | null;

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
        const url = `${this.baseUrl}/home/getClub`;
        return this.http.get<Map<string, Club>>(url);
    }

    likeClub(username: string, clubname: string) {
        const url = `${this.baseUrl}/home/like`;
        const body = { username, clubname };
        return this.http.post(url, body);
    }

    getUsername(): Observable<any> {
        const url = `${this.baseUrl}/home/getUsername`;
        return this.http.get<User>(url);

    }
}