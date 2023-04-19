// shared.service.ts

import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { Club } from './auth.service';

@Injectable({
  providedIn: 'root',
})
export class SharedService {
  private sharedMap = new Subject<Map<string, Club>>();
  sharedMap$ = this.sharedMap.asObservable();

  setSharedMap(map: Map<string, Club>) {
    this.sharedMap.next(map);
  }
}
