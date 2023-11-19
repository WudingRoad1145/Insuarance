package main

var blockSenderContractBytecode string = "0x60806040523480156200001157600080fd5b50604051620027a3380380620027a3833981016040819052620000349162000060565b6000620000428282620001c4565b505062000290565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156200007457600080fd5b82516001600160401b03808211156200008c57600080fd5b818501915085601f830112620000a157600080fd5b815181811115620000b657620000b66200004a565b604051601f8201601f19908116603f01168101908382118183101715620000e157620000e16200004a565b816040528281528886848701011115620000fa57600080fd5b600093505b828410156200011e5784840186015181850187015292850192620000ff565b600086848301015280965050505050505092915050565b600181811c908216806200014a57607f821691505b6020821081036200016b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001bf57600081815260208120601f850160051c810160208610156200019a5750805b601f850160051c820191505b81811015620001bb57828155600101620001a6565b5050505b505050565b81516001600160401b03811115620001e057620001e06200004a565b620001f881620001f1845462000135565b8462000171565b602080601f831160018114620002305760008415620002175750858301515b600019600386901b1c1916600185901b178555620001bb565b600085815260208120601f198616915b82811015620002615788860151825594840194600190910190840162000240565b5085821015620002805787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61250380620002a06000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80637df1cde21161005b5780637df1cde2146100d3578063c2eceb11146100e6578063e884734414610107578063ebb89de41461011a57600080fd5b80633c34db65146100825780634c8820f81461009757806354dfbd39146100c0575b600080fd5b610095610090366004611588565b61012d565b005b6100aa6100a5366004611903565b610193565b6040516100b79190611a4a565b60405180910390f35b6100aa6100ce366004611a5d565b61037f565b6100aa6100e1366004611aae565b61095a565b6100f96100f4366004611903565b6109b2565b6040516100b7929190611b7d565b6100f9610115366004611bab565b610b5d565b6100aa610128366004611a5d565b610c0c565b7f83481d5b04dea534715acad673a8177a46fc93882760f36bdc16ccac439d504e61015b6020830183611cb5565b61016b6040840160208501611cd2565b6101786040850185611cef565b6040516101889493929190611d38565b60405180910390a150565b606061019d610fd0565b6101a657600080fd5b60405163c2eceb1160e01b81526000908190309063c2eceb11906101d4908a908a908a908a90600401611ee0565b600060405180830381865afa1580156101f1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526102199190810190612037565b915091506000806102b46000805461023090612090565b80601f016020809104026020016040519081016040528092919081815260200182805461025c90612090565b80156102a95780601f1061027e576101008083540402835291602001916102a9565b820191906000526020600020905b81548152906001019060200180831161028c57829003601f168201915b505050505084611050565b91509150816102e35730816040516375fff46760e01b81526004016102da9291906120ca565b60405180910390fd5b7f83481d5b04dea534715acad673a8177a46fc93882760f36bdc16ccac439d504e846000015185602001518660400151604051610322939291906120ee565b60405180910390a1604051633c34db6560e01b90610344908690602001612120565b60408051601f19818403018152908290526103629291602001612133565b604051602081830303815290604052945050505050949350505050565b6060610389610fd0565b61039257600080fd5b60006103cb83604051806040016040528060158152602001746d657673686172653a76303a6d617463684269647360581b8152506110f1565b9050600061040e846040518060400160405280601c81526020017f6d657673686172653a76303a756e6d61746368656442756e646c6573000000008152506110f1565b9050805160000361043457306040516375fff46760e01b81526004016102da9190612164565b600081516001600160401b0381111561044f5761044f6115c9565b60405190808252806020026020018201604052801561049c57816020015b6040805160608082018352600080835260208301529181019190915281526020019060019003908161046d5790505b50905060005b82518110156105ef5760008382815181106104bf576104bf612197565b6020026020010151905060005b85518110156105bc57600061052c8783815181106104ec576104ec612197565b602002602001015160000151604051806040016040528060168152602001756d657673686172653a76303a6d65726765644269647360501b8152506111bc565b80602001905181019061053f91906121ad565b90506105828160008151811061055757610557612197565b602002602001015187868151811061057157610571612197565b602002602001015160000151611267565b156105a95786828151811061059957610599612197565b60200260200101519250506105bc565b50806105b481612251565b9150506104cc565b50808383815181106105d0576105d0612197565b60200260200101819052505080806105e790612251565b9150506104a2565b50600081516001600160401b0381111561060b5761060b6115c9565b60405190808252806020026020018201604052801561065057816020015b60408051808201909152600080825260208201528152602001906001900390816106295790505b50905060005b825181101561074e5760006106bd84838151811061067657610676612197565b6020026020010151600001516040518060400160405280601f81526020017f6d657673686172653a76303a65746842756e646c6553696d526573756c7473008152506111bc565b90506000818060200190518101906106d5919061226a565b90506040518060400160405280826001600160401b0316815260200186858151811061070357610703612197565b6020026020010151600001516001600160801b03191681525084848151811061072e5761072e612197565b60200260200101819052505050808061074690612251565b915050610656565b50805160005b61075f600183612287565b81101561086c57600061077382600161229a565b90505b828110156108595783818151811061079057610790612197565b6020026020010151600001516001600160401b03168483815181106107b7576107b7612197565b6020026020010151600001516001600160401b031610156108475760008483815181106107e6576107e6612197565b6020026020010151905084828151811061080257610802612197565b602002602001015185848151811061081c5761081c612197565b60200260200101819052508085838151811061083a5761083a612197565b6020026020010181905250505b8061085181612251565b915050610776565b508061086481612251565b915050610754565b50600083516001600160401b03811115610888576108886115c9565b6040519080825280602002602001820160405280156108b1578160200160208202803683370190505b50905060005b835181101561091b578381815181106108d2576108d2612197565b6020026020010151602001518282815181106108f0576108f0612197565b6001600160801b0319909216602092830291909101909101528061091381612251565b9150506108b7565b5061094b8989836040518060400160405280600b81526020016a06d657673686172653a76360ac1b815250610193565b96505050505050505b92915050565b6060610964610fd0565b61096d57600080fd5b60006109aa8460405180604001604052806019815260200178191959985d5b1d0e9d8c0e989d5a5b19195c94185e5b1bd859603a1b8152506111bc565b949350505050565b604080516060808201835260008083526020830152918101919091526040805160028082526060828101909352600091908160200160208202803683370190505090503081600081518110610a0957610a09612197565b60200260200101906001600160a01b031690816001600160a01b031681525050634210000181600181518110610a4157610a41612197565b60200260200101906001600160a01b031690816001600160a01b0316815250506000610a9b87836040518060400160405280601581526020017464656661756c743a76303a6d65726765644269647360581b81525061132b565b9050610af881600001516040518060400160405280601581526020017464656661756c743a76303a6d65726765644269647360581b81525088604051602001610ae491906122ad565b604051602081830303815290604052611406565b600080610b0a8a8460000151896114b8565b91509150610b4e836000015160405180604001604052806019815260200178191959985d5b1d0e9d8c0e989d5a5b19195c94185e5b1bd859603a1b81525083611406565b50909890975095505050505050565b6040805160608082018352600080835260208301529181019190915260607f67fa9c16cd72410c4cc1d47205b31852a81ec5e92d1c8cebc3ecbe98ed67fe3f846000015184604051610bb09291906122c0565b60405180910390a17f83481d5b04dea534715acad673a8177a46fc93882760f36bdc16ccac439d504e846000015185602001518660400151604051610bf7939291906120ee565b60405180910390a150829050815b9250929050565b6060610c16610fd0565b610c1f57600080fd5b6000610c58836040518060400160405280601581526020017464656661756c743a76303a65746842756e646c657360581b8152506110f1565b90508051600003610c7e57306040516375fff46760e01b81526004016102da9190612164565b600081516001600160401b03811115610c9957610c996115c9565b604051908082528060200260200182016040528015610cde57816020015b6040805180820190915260008082526020820152815260200190600190039081610cb75790505b50905060005b8251811015610ddc576000610d4b848381518110610d0457610d04612197565b6020026020010151600001516040518060400160405280601e81526020017f64656661756c743a76303a65746842756e646c6553696d526573756c747300008152506111bc565b9050600081806020019051810190610d63919061226a565b90506040518060400160405280826001600160401b03168152602001868581518110610d9157610d91612197565b6020026020010151600001516001600160801b031916815250848481518110610dbc57610dbc612197565b602002602001018190525050508080610dd490612251565b915050610ce4565b50805160005b610ded600183612287565b811015610efa576000610e0182600161229a565b90505b82811015610ee757838181518110610e1e57610e1e612197565b6020026020010151600001516001600160401b0316848381518110610e4557610e45612197565b6020026020010151600001516001600160401b03161015610ed5576000848381518110610e7457610e74612197565b60200260200101519050848281518110610e9057610e90612197565b6020026020010151858481518110610eaa57610eaa612197565b602002602001018190525080858381518110610ec857610ec8612197565b6020026020010181905250505b80610edf81612251565b915050610e04565b5080610ef281612251565b915050610de2565b50600083516001600160401b03811115610f1657610f166115c9565b604051908082528060200260200182016040528015610f3f578160200160208202803683370190505b50905060005b8351811015610fa957838181518110610f6057610f60612197565b602002602001015160200151828281518110610f7e57610f7e612197565b6001600160801b03199092166020928302919091019091015280610fa181612251565b915050610f45565b50610fc587878360405180602001604052806000815250610193565b979650505050505050565b6040516000908190819063420100009082818181855afa9150503d8060008114611016576040519150601f19603f3d011682016040523d82523d6000602084013e61101b565b606091505b509150915081611046576342010000816040516375fff46760e01b81526004016102da9291906120ca565b6020015192915050565b6000606061105c610fd0565b61106557600080fd5b60008063421000026001600160a01b031686866040516020016110899291906122e3565b60408051601f19818403018152908290526110a3916122f6565b600060405180830381855afa9150503d80600081146110de576040519150601f19603f3d011682016040523d82523d6000602084013e6110e3565b606091505b509097909650945050505050565b606060008063420300016001600160a01b03168585604051602001611117929190612312565b60408051601f1981840301815290829052611131916122f6565b600060405180830381855afa9150503d806000811461116c576040519150601f19603f3d011682016040523d82523d6000602084013e611171565b606091505b50915091508161119c576342030001816040516375fff46760e01b81526004016102da9291906120ca565b6000818060200190518101906111b29190612334565b9695505050505050565b606060008063420200016001600160a01b031685856040516020016111e29291906122c0565b60408051601f19818403018152908290526111fc916122f6565b600060405180830381855afa9150503d8060008114611237576040519150601f19603f3d011682016040523d82523d6000602084013e61123c565b606091505b5091509150816109aa576342020001816040516375fff46760e01b81526004016102da9291906120ca565b604080516001600160801b03198481166020830152825160108184030181526030830184529084166050830152825180830384018152606090920190925260009190825b825181101561131f578181815181106112c6576112c6612197565b602001015160f81c60f81b6001600160f81b0319168382815181106112ed576112ed612197565b01602001516001600160f81b0319161461130d5760009350505050610954565b8061131781612251565b9150506112ab565b50600195945050505050565b6040805160608082018352600080835260208301529181019190915260008063420300006001600160a01b031686868660405160200161136d939291906123d7565b60408051601f1981840301815290829052611387916122f6565b600060405180830381855afa9150503d80600081146113c2576040519150601f19603f3d011682016040523d82523d6000602084013e6113c7565b606091505b5091509150816113f2576342030000816040516375fff46760e01b81526004016102da9291906120ca565b808060200190518101906111b2919061240b565b60008063420200006001600160a01b031685858560405160200161142c9392919061243f565b60408051601f1981840301815290829052611446916122f6565b600060405180830381855afa9150503d8060008114611481576040519150601f19603f3d011682016040523d82523d6000602084013e611486565b606091505b5091509150816114b1576342020000816040516375fff46760e01b81526004016102da9291906120ca565b5050505050565b60608060008063421000016001600160a01b03168787876040516020016114e193929190612462565b60408051601f19818403018152908290526114fb916122f6565b600060405180830381855afa9150503d8060008114611536576040519150601f19603f3d011682016040523d82523d6000602084013e61153b565b606091505b509150915081611566576342100001816040516375fff46760e01b81526004016102da9291906120ca565b8080602001905181019061157a9190612497565b935093505050935093915050565b60006020828403121561159a57600080fd5b81356001600160401b038111156115b057600080fd5b8201606081850312156115c257600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715611601576116016115c9565b60405290565b60405161010081016001600160401b0381118282101715611601576116016115c9565b604051606081016001600160401b0381118282101715611601576116016115c9565b604051601f8201601f191681016001600160401b0381118282101715611674576116746115c9565b604052919050565b6001600160401b038116811461169157600080fd5b50565b803561169f8161167c565b919050565b60006001600160401b038211156116bd576116bd6115c9565b50601f01601f191660200190565b600082601f8301126116dc57600080fd5b81356116ef6116ea826116a4565b61164c565b81815284602083860101111561170457600080fd5b816020850160208301376000918101602001919091529392505050565b6001600160a01b038116811461169157600080fd5b803561169f81611721565b60006001600160401b0382111561175a5761175a6115c9565b5060051b60200190565b600082601f83011261177557600080fd5b813560206117856116ea83611741565b82815260079290921b840181019181810190868411156117a457600080fd5b8286015b8481101561181b57608081890312156117c15760008081fd5b6117c96115df565b81356117d48161167c565b8152818501356117e38161167c565b818601526040828101356117f681611721565b908201526060828101356118098161167c565b908201528352918301916080016117a8565b509695505050505050565b6000610100828403121561183957600080fd5b611841611607565b905061184c82611694565b815260208201356001600160401b038082111561186857600080fd5b611874858386016116cb565b60208401526040840135604084015261188f60608501611694565b60608401526118a060808501611736565b60808401526118b160a08501611694565b60a084015260c084013560c084015260e08401359150808211156118d457600080fd5b506118e184828501611764565b60e08301525092915050565b6001600160801b03198116811461169157600080fd5b6000806000806080858703121561191957600080fd5b84356001600160401b038082111561193057600080fd5b61193c88838901611826565b9550602091508187013561194f8161167c565b945060408701358181111561196357600080fd5b8701601f8101891361197457600080fd5b80356119826116ea82611741565b81815260059190911b8201840190848101908b8311156119a157600080fd5b928501925b828410156119c85783356119b9816118ed565b825292850192908501906119a6565b965050505060608701359150808211156119e157600080fd5b506119ee878288016116cb565b91505092959194509250565b60005b83811015611a155781810151838201526020016119fd565b50506000910152565b60008151808452611a368160208601602086016119fa565b601f01601f19169290920160200192915050565b6020815260006115c26020830184611a1e565b60008060408385031215611a7057600080fd5b82356001600160401b03811115611a8657600080fd5b611a9285828601611826565b9250506020830135611aa38161167c565b809150509250929050565b60008060408385031215611ac157600080fd5b8235611acc816118ed565b915060208301356001600160401b03811115611ae757600080fd5b611af3858286016116cb565b9150509250929050565b600081518084526020808501945080840160005b83811015611b365781516001600160a01b031687529582019590820190600101611b11565b509495945050505050565b6001600160801b031981511682526001600160401b03602082015116602083015260006040820151606060408501526109aa6060850182611afd565b604081526000611b906040830185611b41565b8281036020840152611ba28185611a1e565b95945050505050565b60008060408385031215611bbe57600080fd5b82356001600160401b0380821115611bd557600080fd5b9084019060608287031215611be957600080fd5b611bf161162a565b8235611bfc816118ed565b8152602083810135611c0d8161167c565b82820152604084013583811115611c2357600080fd5b80850194505087601f850112611c3857600080fd5b8335611c466116ea82611741565b81815260059190911b8501820190828101908a831115611c6557600080fd5b958301955b82871015611c8c578635611c7d81611721565b82529583019590830190611c6a565b60408501525091955086013592505080821115611ca857600080fd5b50611af3858286016116cb565b600060208284031215611cc757600080fd5b81356115c2816118ed565b600060208284031215611ce457600080fd5b81356115c28161167c565b6000808335601e19843603018112611d0657600080fd5b8301803591506001600160401b03821115611d2057600080fd5b6020019150600581901b3603821315610c0557600080fd5b6000606082016001600160801b03198716835260206001600160401b03871681850152606060408501528185835260808501905086925060005b86811015611da0578335611d8581611721565b6001600160a01b031682529282019290820190600101611d72565b5098975050505050505050565b600081518084526020808501945080840160005b83811015611b3657815180516001600160401b039081168952848201518116858a01526040808301516001600160a01b0316908a0152606091820151169088015260809096019590820190600101611dc1565b60006101006001600160401b038084511685526020840151826020870152611e3e83870182611a1e565b925050604084015160408601528060608501511660608601525060018060a01b03608084015116608085015260a0830151611e8460a08601826001600160401b03169052565b5060c083015160c085015260e083015184820360e0860152611ba28282611dad565b600081518084526020808501945080840160005b83811015611b365781516001600160801b03191687529582019590820190600101611eba565b608081526000611ef36080830187611e14565b6001600160401b03861660208401528281036040840152611f148186611ea6565b90508281036060840152610fc58185611a1e565b600060608284031215611f3a57600080fd5b611f4261162a565b90508151611f4f816118ed565b8152602082810151611f608161167c565b8282015260408301516001600160401b03811115611f7d57600080fd5b8301601f81018513611f8e57600080fd5b8051611f9c6116ea82611741565b81815260059190911b82018301908381019087831115611fbb57600080fd5b928401925b82841015611fe2578351611fd381611721565b82529284019290840190611fc0565b6040860152509295945050505050565b600082601f83011261200357600080fd5b81516120116116ea826116a4565b81815284602083860101111561202657600080fd5b6109aa8260208301602087016119fa565b6000806040838503121561204a57600080fd5b82516001600160401b038082111561206157600080fd5b61206d86838701611f28565b9350602085015191508082111561208357600080fd5b50611af385828601611ff2565b600181811c908216806120a457607f821691505b6020821081036120c457634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160a01b03831681526040602082018190526000906109aa90830184611a1e565b6001600160801b0319841681526001600160401b0383166020820152606060408201526000611ba26060830184611afd565b6020815260006115c26020830184611b41565b6001600160e01b03198316815281516000906121568160048501602087016119fa565b919091016004019392505050565b6001600160a01b03919091168152604060208201819052600790820152666e6f206269647360c81b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600060208083850312156121c057600080fd5b82516001600160401b038111156121d657600080fd5b8301601f810185136121e757600080fd5b80516121f56116ea82611741565b81815260059190911b8201830190838101908783111561221457600080fd5b928401925b82841015610fc557835161222c816118ed565b82529284019290840190612219565b634e487b7160e01b600052601160045260246000fd5b6000600182016122635761226361223b565b5060010190565b60006020828403121561227c57600080fd5b81516115c28161167c565b818103818111156109545761095461223b565b808201808211156109545761095461223b565b6020815260006115c26020830184611ea6565b6001600160801b0319831681526040602082015260006109aa6040830184611a1e565b604081526000611b906040830185611a1e565b600082516123088184602087016119fa565b9190910192915050565b6001600160401b03831681526040602082015260006109aa6040830184611a1e565b6000602080838503121561234757600080fd5b82516001600160401b038082111561235e57600080fd5b818501915085601f83011261237257600080fd5b81516123806116ea82611741565b81815260059190911b8301840190848101908883111561239f57600080fd5b8585015b83811015611da0578051858111156123bb5760008081fd5b6123c98b89838a0101611f28565b8452509186019186016123a3565b6001600160401b03841681526060602082015260006123f96060830185611afd565b82810360408401526111b28185611a1e565b60006020828403121561241d57600080fd5b81516001600160401b0381111561243357600080fd5b6109aa84828501611f28565b6001600160801b0319841681526060602082015260006123f96060830185611a1e565b6060815260006124756060830186611e14565b6001600160801b03198516602084015282810360408401526111b28185611a1e565b600080604083850312156124aa57600080fd5b82516001600160401b03808211156124c157600080fd5b61206d86838701611ff256fea264697066735822122076e3f4ec81afd8028b2fdc226aa6a72913298545ae3065d1dd95c2c0f12d482064736f6c63430008130033"
