import { Injectable } from '@angular/core';
import { Http, Response, URLSearchParams } from '@angular/http';
import { GlobalData } from '../interfaces/global-data';
import { SocketService } from './socket.service';
import { UserService } from './user.service';
import { KenTheGuruService } from './ken-the-guru.service';
import { KmiService } from './kmi.service';
import { user, kentheguru, kmi } from '../../messages/messages';
import { ProtoResponse } from '../interfaces/proto-response';
import { Observable, Subject } from 'rxjs/Rx';

@Injectable()
export class GlobalDataService {
  private gd: GlobalData;

  constructor(private http: Http, private us: UserService, private ktgs: KenTheGuruService, private kmis: KmiService) {
    this.gd = {};
  }
  
  /* User methods */

  getUserId(): number {
    if(this.gd.user && this.gd.user.ID) {
      return this.gd.user.ID;
    } else {
      return -1;
    }
  }

  getUserSnapshot(): user.User {
    if(this.gd.user) {
      return this.gd.user;
    } else {
      return undefined;
    }
  }

  setAndGetUserById(id: number): Observable<user.User> {
    let obs = this.us
      .reconnect()
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'GetUserResponse';
      })
      .map((value: ProtoResponse): user.User => {
        return user.User.create(user.GetUserResponse.from(value.data).user);
      });

    obs.subscribe(
      user => {
        this.gd.user = user;
      },
      error => {
        console.log(error);
      }
    );
    
    this.us.next('GetUserRequest', {
      ID: id
    });

    return obs;
  }

  registerUser(userProperties: user.User$Properties): Observable<number> {
    let obs = this.us
      .reconnect()
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'CreateUserResponse';
      })
      .map((value: ProtoResponse): number => {
        return user.CreateUserResponse.from(value.data).ID;
      });

    this.us.next('CreateUserRequest', user);

    return obs;
  }

  editUser(id: number, configProperties: user.Config$Properties): Observable<boolean> {
    let obs = this.us
      .reconnect()
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'EditUserResponse';
      })
      .map((value: ProtoResponse): boolean => {
        return user.EditUserResponse.from(value.data).error ? false : true;
      });

    obs.subscribe(
      success => {
        if(success) {
          this.setAndGetUserById(this.getUserId());
        }
      },
      error => {
        console.log(error);
      }
    )

    this.us.next('EditUserRequest', {
      ID: id,
      config: configProperties
    });

    return obs;
  }

  logIn(username: string, password: string): Subject<Response> {
    let obs = this.ktgs
      .reconnect()
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'AuthenticationResponse' || value.message == 'ErrorResponse';
      })
      .map((value: ProtoResponse): string => {
        if(value.message == 'AuthenticationResponse') {
          return kentheguru.AuthenticationResponse.from(value.data).token;
        } else {
          return 'nope';
        }
      });

    let cookieRequest = new Subject<Response>();

    obs.subscribe(
      token => {
        let body = new URLSearchParams();
        body.set('token', token);

        this.http.post(`http://${SocketService.SOCKET_ADDRESS}/auth`, body).share().subscribe(
          data => {
            if(data.text() == 'Authenticated!') {
              cookieRequest.next(data);
            }
          },
          error => {
            console.log(error);
          }
        );
      },
      error => {
        console.log(error);
      }
    );

    this.ktgs.next('AuthenticationRequest', {
      username: username,
      password: password
    });

    return cookieRequest;
  }

  /* KMI methods */

  addKMI(path: string): Observable<number> {
    let obs = this.kmis
      .reconnect()
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'AddKMIResponse';
      })
      .map((value: ProtoResponse): number => {
        let akr = kmi.AddKMIResponse.from(value.data);
        if(!akr.error) {
          return akr.ID;
        }
    });

    this.kmis.next('AddKMIRequest', {
      path: path
    });

    return obs;
  }
}
