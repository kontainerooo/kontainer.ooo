import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'kio-avatar',
  templateUrl: './avatar.component.html',
  styleUrls: ['./avatar.component.scss']
})
export class AvatarComponent implements OnInit {
  @Input('src') avatarImage: string;

  constructor() { }

  ngOnInit() {
  }

}
