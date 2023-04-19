// shared.service.ts

import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { Club } from './auth.service';

@Injectable({
  providedIn: 'root',
})
export class SharedService {
  private sharedMap = new Subject<Map<string, Club>>();
  private selectedPair = new Subject<[string, Club]>;

  setSharedMap(map: Map<string, Club>) {
    this.sharedMap.next(map);
  }

  getSharedMap(){
    return this.sharedMap.asObservable();
  }

  setSelectedPair(key: string, value: Club) {
    this.selectedPair.next([key, value]);
  }

  getSelectedPair() {
    return this.selectedPair.asObservable();
  }
}
