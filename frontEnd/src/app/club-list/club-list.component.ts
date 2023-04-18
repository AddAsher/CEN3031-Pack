import { AuthService, Club } from '../auth.service';
import { Component, EventEmitter,Output } from '@angular/core';

@Component({
  selector: 'app-club-list',
  templateUrl: './club-list.component.html',
  styleUrls: ['./club-list.component.css']
})
export class ClubListComponent {
  clubs: Map<string, Club> = new Map<string, Club>();

  @Output() sender = new EventEmitter();

  constructor(private authService: AuthService) { }

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
    
    this.sender.emit(this.clubs);
  }

}