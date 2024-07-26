import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class CacheService {

  private baseUrl = 'http://localhost:8080'

  constructor(private http: HttpClient) { }

  addValue(key: string, value: string, expiry: number) {
    return this.http.post(`${this.baseUrl}/cache`, { key, value, expiration: `${expiry}s` })
  }

  getCacheList() {
    return this.http.get(`${this.baseUrl}/cache`)
  }

  deleteValue(key: string) {
    return this.http.delete(`${this.baseUrl}/cache/${key}`)
  }
}
