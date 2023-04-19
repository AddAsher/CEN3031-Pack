// shared.service.ts

import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { NullableClub } from './auth.service';

@Injectable({
  providedIn: 'root',
})
export class SharedService {
  private sharedMap = new BehaviorSubject<Map<string, NullableClub>>(new Map<string, NullableClub>());
  private selectedPair = new BehaviorSubject<[string, NullableClub]>(['', null]);

  setSharedMap(map: Map<string, NullableClub>) {
    this.sharedMap.next(map);
  }

  getSharedMap(){
    return this.sharedMap.asObservable();
  }

  setSelectedPair(key: string, value: NullableClub) {
    this.selectedPair.next([key, value]);
  }

  getSelectedPair() {
    return this.selectedPair.asObservable();
  }
}
