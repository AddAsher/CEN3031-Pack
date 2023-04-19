import { Component} from '@angular/core';
import { SharedService } from '../shared.service';
import { Club } from '../auth.service';

@Component({
  selector: 'app-club-box',
  templateUrl: './club-box.component.html',
  styleUrls: ['./club-box.component.css']
})
export class ClubBoxComponent {
  club: [string,any];

  constructor(private sharedService: SharedService){
    this.club=["",null];
  }

  ngOnInit(){
    this.sharedService.getSelectedPair().subscribe(pair => {
      this.club = pair;
    });
  }
}
