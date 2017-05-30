import { Component, OnInit, OnDestroy } from '@angular/core';
import { KenTheGuruService } from '../services/ken-the-guru.service';
import {
  FormGroup,
  Validators,
  FormBuilder
} from '@angular/forms';

@Component({
  selector: 'kro-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.scss']
})
export class SignInComponent implements OnInit, OnDestroy {
  form: FormGroup;

  constructor(private fb: FormBuilder, private ktgs: KenTheGuruService) {
    this.form = this.fb.group({
      username: [''],
      password: ['']
    });
  }

  ngOnInit() {
    this.ktgs.messages.subscribe(
      value => {
        console.log(value);
      },
      error => {
        console.log(error);
      }
    );
  }

  ngOnDestroy() {
    this.ktgs.messages.unsubscribe();
  }

  onSubmit() {
    this.ktgs.next('AuthenticationRequest', {
      username: this.form.get('username').value,
      password: this.form.get('password').value
    });
  }

}
