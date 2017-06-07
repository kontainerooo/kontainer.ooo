import { Injectable } from '@angular/core';
import { Http, Response, URLSearchParams } from '@angular/http';
import { GlobalData } from '../interfaces/global-data';
import { SocketService } from './socket.service';
import { UserService } from './user.service';
import { KenTheGuruService } from './ken-the-guru.service';
import { KmiService } from './kmi.service';
import { ContainerService } from './container.service';
import { ModuleService } from './module.service';

import { user, kentheguru, kmi, container, module } from '../../messages/messages';
import { ProtoResponse } from '../interfaces/proto-response';
import { Observable, Subject } from 'rxjs/Rx';

@Injectable()
export class GlobalDataService {
  private gd: GlobalData;

  constructor(
    private http: Http,
    private us: UserService,
    private ktgs: KenTheGuruService,
    private kmis: KmiService,
    private cs: ContainerService,
    private ms: ModuleService
  ) {
    this.gd = {};
  }
  
  /* User methods */

  getUserIdSnapshot(): number {
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

  setAndGetUserAutomatically(): Observable<user.User> {
    if(this.getUserIdSnapshot() !== -1) {
      return this.setAndGetUserById(this.getUserIdSnapshot());
    }
    return this.setAndGetUserByCookie();
  }

  setAndGetUserByCookie(): Observable<user.User> {
    return this.setAndGetUser();
  }

  setAndGetUserById(id: number): Observable<user.User> {
    return this.setAndGetUser(id);
  }

  private setAndGetUser(id?: number): Observable<user.User> {
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

    let requestObject = {
      ID: 0
    };
    if(id != undefined || id != null) {
      requestObject = {
        ID: id
      };
    }
    
    this.us.next('GetUserRequest', requestObject);

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
          this.setAndGetUserById(this.getUserIdSnapshot());
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
            this.setAndGetUserById(1);
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

  setAndGetAvailableKMI(): Observable<kmi.KMDI[]> {
    let obs = this.kmis
      .reconnect()
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'KMIResponse';
      })
      .map((value: ProtoResponse): kmi.KMDI[] => {
        let kr = kmi.KMIResponse.from(value.data);
        let kmdiArray: kmi.KMDI[] = [];
        if(!kr.error) {
          for(let i in kr.kmdi) {
            kr.kmdi[i].type = (<any>value.data).kmdi[i].type;
            kmdiArray.push(kmi.KMDI.from(kr.kmdi[i]));
          }
          return kmdiArray;
        }
      });

    obs.subscribe(
      value => {
        this.gd.KMDI = value;
      },
      error => {
        console.log(error);
      }
    )

    this.kmis.next('KMIRequest', {});

    return obs;
  }

  /* Container methods */

  createContainer(kmiId: number, name: string): Observable<string> {
    let obs = this.cs
      .reconnect()
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'CreateContainerResponse';
      })
      .map((value: ProtoResponse): string => {
        let ccr = container.CreateContainerResponse.from(value.data);
        console.log(ccr);
        if(!ccr.error) {
          return ccr.ID;
        }
      });

    this.cs.next('CreateContainerRequest', {
      refID: this.getUserIdSnapshot(),
      kmiID: kmiId,
      name: name
    });

    return obs;
  }

  getContainers(): Observable<container.container[]> {
    let obs = this.cs
      .reconnect()
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'InstancesResponse';
      })
      .map((value: ProtoResponse): container.container[] => {
        let ir = container.InstancesResponse.from(value.data);
        let containerArray: container.container[] = [];
        for(let i in ir.instances) {
          ir.instances[i].kmi.KMDI.type = (<any>value.data).instances[i].kmi.KMDI.type;
          containerArray.push(container.container.from(ir.instances[i]));
        }
        return containerArray;
      });
    
    this.cs.next('InstancesRequest', {
      refID: this.getUserIdSnapshot()
    });

    return obs;
  }

  removeContainer(refId: number, id: string): Observable<boolean> {
    let obs = this.cs
      .reconnect()
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'RemoveContainerResponse';
      })
      .map((value: ProtoResponse): boolean => {
        let rcr = container.RemoveContainerResponse.from(value.data);
        if(!rcr.error) {
          return true;
        }
      });

    this.cs.next('RemoveContainerRequest', {
      refID: refId,
      ID: id
    });

    return obs;
  }

  /* Module methods */

  setAndGetContainerKMI(name: string, containerId: string): Observable<kmi.KMI> {
    let obs = this.cs
      .reconnect()
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'GetContainerKMIResponse';
      })
      .map((value: ProtoResponse): kmi.KMI => {
        let gmcr = container.GetContainerKMIResponse.from(value.data);
        if(!gmcr.error) {
          gmcr.containerKMI.KMDI.type = (<any>value.data).containerKMI.KMDI.type;
          return kmi.KMI.from(gmcr.containerKMI);
        }
      });

    obs.subscribe(
      value => {
        this.gd.currentKMI = {
          name: name,
          kmi: value
        };
      },
      error => {
        console.log(error);
      }
    )

    this.cs.next('GetContainerKMIRequest', {
      containerID: containerId
    });

    return obs;
  }

  getValueSnapshot(template: string, name: string): string {
    console.log(this.gd.currentKMI);
    for(let module of this.gd.currentKMI.kmi.frontend) {
      if(module.template == template) {
        return module.parameters[name];
      }
    }
  }

  sendCommand(command: string, env?: {[k: string]: string}) {
    if(!env) {
      let env = {};
    }

    let obs = this.ms
      .reconnect()
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'SendCommandResponse';
      })
      .map((value: ProtoResponse): string => {
        let scr = module.SendCommandResponse.from(value.data);
        if(!scr.error) {
          return scr.response;
        }
      });

    this.ms.next('SendCommandRequest', {
      refID: this.getUserIdSnapshot(),
      containerName: this.gd.currentKMI.name,
      command: command,
      env: env
    });

    return obs;
  }
}
