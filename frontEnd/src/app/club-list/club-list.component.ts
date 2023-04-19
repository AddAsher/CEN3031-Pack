import { AuthService, Club } from '../auth.service';
import { Component } from '@angular/core';
import { SharedService } from '../shared.service';

@Component({
  selector: 'app-club-list',
  templateUrl: './club-list.component.html',
  styleUrls: ['./club-list.component.css']
})
export class ClubListComponent {
  clubs: Map<string, Club> = new Map<string, Club>();

  constructor(private authService: AuthService, private sharedService: SharedService) { }

  ngOnInit() {
    this.authService.getClubs().subscribe(
      (response: Map<string, Club>) => {
        this.clubs = response;
        console.log(response);
      },
      (error) => {
        console.log(error);
      }
    );
    this.sharedService.setSharedMap(this.clubs);
  }

}