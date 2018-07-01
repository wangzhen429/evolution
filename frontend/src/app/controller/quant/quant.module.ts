import { NgModule }                  from '@angular/core';
import { BrowserModule }             from '@angular/platform-browser';
import { BrowserAnimationsModule   } from '@angular/platform-browser/animations';
import { HttpClientModule  }         from '@angular/common/http';
import { HttpModule   }              from '@angular/http';
import { FormsModule }               from '@angular/forms';
import { ClarityModule  }            from "@clr/angular";
import { TranslateModule  }          from "@ngx-translate/core";

import { PoolComponent }         from './pool/pool.component';
import { PoolAddItemComponent }  from './pool/add-item/pool-add-item.component';
import { PoolListItemComponent } from './pool/list-item/pool-list-item.component';
import { SavePoolComponent }     from './pool/save/save.component';
import { ClassifyComponent }     from './classify/classify.component';
import { ItemComponent }         from './item/item.component';
import { AssetSourceComponent }  from './config/asset-source/asset-source.component';

import { PoolService }            from '../../service/quant/pool.service';
import { ConfigService }          from '../../service/quant/config.service';
import { ClassifyService }        from '../../service/quant/classify.service';
import { ItemService }            from '../../service/quant/item.service';
import { MessageHandlerService  } from '../../service/base/message-handler.service';
import { SocketService }          from '../../service/base/socket.service';
import { WebsocketService }       from '../../service/base/websocket.service';


@NgModule({
  declarations: [
    PoolComponent,
    PoolAddItemComponent,
    PoolListItemComponent,
    SavePoolComponent,
    ClassifyComponent,
    ItemComponent,
    AssetSourceComponent,
  ],
  imports: [   
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    HttpModule,
    FormsModule,
    ClarityModule,
    TranslateModule,
  ],
  providers: [ 
    PoolService,
    ConfigService,
    ClassifyService,
    ItemService,
    MessageHandlerService,
    SocketService,
    WebsocketService,
  ],
  exports: [
  ]
})

export class QuantModule {
}
