import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router } from '@angular/router';
import { GlobalDataService } from '../services/global-data.service';
import { MdlSnackbarService } from 'angular2-mdl';
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

  constructor(private fb: FormBuilder, private gds: GlobalDataService, private mdlss: MdlSnackbarService, private router: Router) {
    this.form = this.fb.group({
      username: [''],
      password: ['']
    });
  }

  ngOnInit() {

  }

  ngOnDestroy() {

  }

  onSubmit() {
    this.mdlss.showToast('Logging in...');
    this.gds.logIn(this.form.get('username').value, this.form.get('password').value).subscribe(
      value => {
        if(value) {
          this.router.navigate(['/']);
        }
      },
      error => {
        console.log(error);
      }
    );
  }

}
