import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClubBoxComponent } from './club-box.component';

describe('ClubBoxComponent', () => {
  let component: ClubBoxComponent;
  let fixture: ComponentFixture<ClubBoxComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ClubBoxComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ClubBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
