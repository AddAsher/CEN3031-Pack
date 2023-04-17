import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddClubComponent } from './add-club.component';

describe('AddClubComponent', () => {
  let component: AddClubComponent;
  let fixture: ComponentFixture<AddClubComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddClubComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddClubComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
