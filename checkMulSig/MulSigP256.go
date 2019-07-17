// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package checkMulSig

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MulSigP256ABI is the input ABI used to generate the binding from.
const MulSigP256ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"}],\"name\":\"twice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"multipleGeneratorByScalar\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"isOnCurve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zeroProj\",\"outputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"},{\"name\":\"exp\",\"type\":\"uint256\"}],\"name\":\"multiplyPowerBase2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"},{\"name\":\"z0\",\"type\":\"uint256\"}],\"name\":\"toAffinePoint\",\"outputs\":[{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"},{\"name\":\"x2\",\"type\":\"uint256\"},{\"name\":\"y2\",\"type\":\"uint256\"}],\"name\":\"addAndReturnProjectivePoint\",\"outputs\":[{\"name\":\"P\",\"type\":\"uint256[3]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"multiplyScalar\",\"outputs\":[{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"}],\"name\":\"toProjectivePoint\",\"outputs\":[{\"name\":\"P\",\"type\":\"uint256[3]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"xPks\",\"type\":\"uint256[8]\"},{\"name\":\"yPks\",\"type\":\"uint256[8]\"},{\"name\":\"idxSigs\",\"type\":\"uint256[8]\"},{\"name\":\"sigLen\",\"type\":\"uint8\"},{\"name\":\"bytesR\",\"type\":\"bytes\"},{\"name\":\"mess\",\"type\":\"bytes32\"}],\"name\":\"genCommonParams\",\"outputs\":[{\"name\":\"res\",\"type\":\"uint256[3]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zeroAffine\",\"outputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"xPks\",\"type\":\"uint256[8]\"},{\"name\":\"yPks\",\"type\":\"uint256[8]\"},{\"name\":\"idxSigs\",\"type\":\"uint256[8]\"},{\"name\":\"sigsLen\",\"type\":\"int128\"},{\"name\":\"xR\",\"type\":\"uint256\"},{\"name\":\"yR\",\"type\":\"uint256\"},{\"name\":\"bytesR\",\"type\":\"bytes\"},{\"name\":\"sig\",\"type\":\"uint256\"},{\"name\":\"mess\",\"type\":\"bytes32\"}],\"name\":\"checkMulSig\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"}],\"name\":\"isZeroCurve\",\"outputs\":[{\"name\":\"isZero\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"z\",\"type\":\"uint256\"}],\"name\":\"twiceProj\",\"outputs\":[{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"},{\"name\":\"z1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"num\",\"type\":\"uint256\"},{\"name\":\"modular\",\"type\":\"uint256\"}],\"name\":\"inverseMod\",\"outputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"},{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x0\",\"type\":\"uint256\"},{\"name\":\"y0\",\"type\":\"uint256\"},{\"name\":\"z0\",\"type\":\"uint256\"},{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"},{\"name\":\"z1\",\"type\":\"uint256\"}],\"name\":\"addProj\",\"outputs\":[{\"name\":\"x2\",\"type\":\"uint256\"},{\"name\":\"y2\",\"type\":\"uint256\"},{\"name\":\"z2\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MulSigP256Bin is the compiled bytecode used for deploying new contracts.
const MulSigP256Bin = `0x608060405234801561001057600080fd5b5061152f806100206000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80637ec8da8d116100a2578063c30cfa2d11610071578063c30cfa2d14610563578063c80edca414610586578063ca465e6e146105af578063e022d77c146105e4578063f214aba0146106135761010b565b80637ec8da8d146102a857806380a694de146102cb57806384dfba461461040c578063883c48d4146104145761010b565b8063322b24aa116100de578063322b24aa146101c6578063675ca043146101ef578063713eca281461021857806372fb4a141461027f5761010b565b806309d3ef31146101105780630afb4ddc1461014c5780630b0dbcfa1461016957806314c67060146101a0575b600080fd5b6101336004803603604081101561012657600080fd5b508035906020013561064e565b6040805192835260208301919091528051918290030190f35b6101336004803603602081101561016257600080fd5b503561067d565b61018c6004803603604081101561017f57600080fd5b50803590602001356106d4565b604080519115158252519081900360200190f35b6101a86107c6565b60408051938452602084019290925282820152519081900360600190f35b610133600480360360608110156101dc57600080fd5b50803590602081013590604001356107d0565b6101336004803603606081101561020557600080fd5b5080359060208101359060400135610817565b6102476004803603608081101561022e57600080fd5b5080359060208101359060408101359060600135610867565b6040518082606080838360005b8381101561026c578181015183820152602001610254565b5050505090500191505060405180910390f35b6101336004803603606081101561029557600080fd5b5080359060208101359060400135610898565b610247600480360360408110156102be57600080fd5b5080359060200135610966565b61024760048036036103608110156102e257600080fd5b810190808061010001906008806020026040519081016040528092919082600860200280828437600092019190915250506040805161010081810190925292959493818101939250906008908390839080828437600092019190915250506040805161010081810190925292959493818101939250906008908390839080828437600092019190915250919460ff8435169490939092506040810191506020013564010000000081111561039557600080fd5b8201836020820111156103a757600080fd5b803590602001918460018302840111640100000000831117156103c957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506109bf915050565b610133610d8e565b61018c60048036036103c081101561042b57600080fd5b81019080806101000190600880602002604051908101604052809291908260086020028082843760009201919091525050604080516101008181019092529295949381810193925090600890839083908082843760009201919091525050604080516101008181019092529295949381810193925090600890839083908082843760009201919091525091948335600f0b9460208501359460408101359450919250906080810190606001356401000000008111156104e957600080fd5b8201836020820111156104fb57600080fd5b8035906020019184600183028401116401000000008311171561051d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610d95565b61018c6004803603604081101561057957600080fd5b5080359060200135610e6a565b6101a86004803603606081101561059c57600080fd5b5080359060208101359060400135610e8e565b6105d2600480360360408110156105c557600080fd5b50803590602001356110d6565b60408051918252519081900360200190f35b610133600480360360808110156105fa57600080fd5b5080359060208101359060408101359060600135611146565b6101a8600480360360c081101561062957600080fd5b5080359060208101359060408101359060608101359060808101359060a0013561117b565b600080600061065f85856001610e8e565b91965094509050610671858583610817565b92509250509250929050565b6000806106cb7f6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c2967f4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f585610898565b91509150915091565b60008215806106f05750600160601b63ffffffff60c01b031983145b806106f9575081155b806107115750600160601b63ffffffff60c01b031982145b1561071e575060006107c0565b6000600160601b63ffffffff60c01b031983840990506000600160601b63ffffffff60c01b031985600160601b63ffffffff60c01b0319878809099050600160601b63ffffffff60c01b0319806bfffffffffffffffffffffffd63ffffffff60c01b0319870982089050600160601b63ffffffff60c01b03197f5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b820890501490505b92915050565b6000600181909192565b60008084846001835b868110156107fb576107ec848484610e8e565b919550935091506001016107d9565b50610807838383610817565b945094505050505b935093915050565b600080600061083484600160601b63ffffffff60c01b03196110d6565b9050600160601b63ffffffff60c01b03198187099250600160601b63ffffffff60c01b0319818609915050935093915050565b61086f611482565b60008061087e87878787611146565b909250905061088d8282610966565b979650505050505050565b600080848484806108b7576108ab610d8e565b9450945050505061080f565b80600114156108cb5750909250905061080f565b80600214156108de576108ab838361064e565b91935091508290829082826001808481166108fb57600097508798505b600185901c94505b841561094857610914848484610e8e565b919550935091506001808616141561093c576109348484848c8c8661117b565b919a50985090505b600185901c9450610903565b610953898983610817565b9850985050505050505050935093915050565b61096e611482565b600160601b63ffffffff60c01b0319600160000860408201819052600160601b63ffffffff60c01b031990840981526040810151600160601b63ffffffff60c01b0319908309602082015292915050565b6109c7611482565b6109cf6114a0565b6109d76114be565b6109df6114dc565b895183528851825260018082526060905b6008811015610a68578a8160088110610a0557fe5b60200201518c8260088110610a1657fe5b602002015117610a2557610a68565b845184518451610a5a9291908f8560088110610a3d57fe5b60200201518f8660088110610a4e57fe5b6020020151600161117b565b8552855285526001016109f0565b508351610a86908460005b60200201518460005b6020020151610817565b84528452600060208086018290528481018290526001848201819052845260408087018390528086018390526080870183905282885290870182905286018190525b6008811015610c4e578a8160088110610add57fe5b60200201518c8260088110610aee57fe5b602002015117610afd57610c4e565b84518451610b2d91908e8460088110610b1257fe5b60200201518e8560088110610b2357fe5b6020020151611146565b6060868101919091528601819052610b4d908560035b6020020151611280565b805160208201207bffffffff00000000000000004319055258e8617b0c46353d039cdaae19900660808701529150610bae8c8260088110610b8a57fe5b60200201518c8360088110610b9b57fe5b60200201518760045b6020020151610898565b60408681019190915286015285518a9060088110610bc857fe5b602002015181148015610bde5750855160ff8a16115b15610c1c5760208601516040870151610c0891908560005b60200201516040890151886002610a4e565b855260408801526020870152855160010186525b60208086015190850151610c339190856001610bf6565b60208681019190915286810191909152860152600101610ac8565b5060208085015190840151610c669190846001610a7c565b6060858101919091528501819052610c8090846003610b43565b90508087876040516020018084805190602001908083835b60208310610cb75780518252601f199092019160209182019101610c98565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310610cff5780518252601f199092019160209182019101610ce0565b51815160209384036101000a6000190180199092169116179052920193845250604080518085038152938201905282519201919091208089527bffffffff00000000000000004319055258e8617b0c46353d039cdaae1993509150610d619050565b0685526020850151610d7590866002610a73565b6040870152602086015250929998505050505050505050565b6000809091565b6000600887600f0b1315610dab57506000610e5d565b86610db4611482565b610dc28c8c8c858a896109bf565b90506000808080610e147f6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c2967f4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f58b610898565b60208701516040880151929650909450610e3091876000610ba4565b9092509050610e4182828e8e611146565b90925090508382148015610e5457508083145b96505050505050505b9998505050505050505050565b600082158015610e78575081155b15610e85575060016107c0565b50600092915050565b6000808085858583808080610ea38787610e6a565b15610ec257610eb06107c6565b995099509950505050505050506110cd565b600160601b63ffffffff60c01b03198587099250600160601b63ffffffff60c01b0319600284099250600160601b63ffffffff60c01b03198784099150600160601b63ffffffff60c01b03198683099150600160601b63ffffffff60c01b0319600283099150600160601b63ffffffff60c01b03198788099650600160601b63ffffffff60c01b0319600388099350600160601b63ffffffff60c01b03198586099450600160601b63ffffffff60c01b03196bfffffffffffffffffffffffd63ffffffff60c01b031986099450600160601b63ffffffff60c01b03198585089350600160601b63ffffffff60c01b03198485099050600160601b63ffffffff60c01b0319826002099650600160601b63ffffffff60c01b031987600160601b63ffffffff60c01b03190382089050600160601b63ffffffff60c01b031981600160601b63ffffffff60c01b03190383089650600160601b63ffffffff60c01b03198785099650600160601b63ffffffff60c01b03198387099550600160601b63ffffffff60c01b03198687099550600160601b63ffffffff60c01b0319866002099550600160601b63ffffffff60c01b031986600160601b63ffffffff60c01b03190388089850600160601b63ffffffff60c01b03198184099950600160601b63ffffffff60c01b03198384099750600160601b63ffffffff60c01b03198389099750505050505050505b93509350939050565b60008215806110e457508183145b156110f1575060006107c0565b6000808385048486068560015b60018211156111205794859450828204935090829006919084840290036110fe565b5050505050600081121561113b57806000038303915061113f565b8091505b5092915050565b600080600061115b878760018888600161117b565b9198509650905061116d878783610817565b925092505094509492505050565b60008060008060008060006111908d8d610e6a565b156111a75789898996509650965050505050611274565b6111b18a8a610e6a565b156111c8578c8c8c96509650965050505050611274565b600160601b63ffffffff60c01b0319888d099350600160601b63ffffffff60c01b03198b8a099250600160601b63ffffffff60c01b0319888e099150600160601b63ffffffff60c01b03198b8b0990508082141561124a5782841415611242576112338d8d8d610e8e565b96509650965050505050611274565b6112336107c6565b611268600160601b63ffffffff60c01b0319898d09838386886112cf565b91985096509450505050505b96509650969350505050565b6060600160f91b6001808416141561129957600160f81b175b604080516001600160f81b0319909216602083015260218083018690528151808403909101815260419092019052905092915050565b6000808087878787878580808080600160601b63ffffffff60c01b031987600160601b63ffffffff60c01b03190387089050600160601b63ffffffff60c01b031988600160601b63ffffffff60c01b0319038a089450600160601b63ffffffff60c01b03198586099350600160601b63ffffffff60c01b03198182099150600160601b63ffffffff60c01b03198a83099150600160601b63ffffffff60c01b03198989089750600160601b63ffffffff60c01b03198489099750600160601b63ffffffff60c01b031988600160601b63ffffffff60c01b03190383089150600160601b63ffffffff60c01b03198286099c50600160601b63ffffffff60c01b03198585099250600160601b63ffffffff60c01b0319848a099850600160601b63ffffffff60c01b031982600160601b63ffffffff60c01b0319038a089850600160601b63ffffffff60c01b03198982099050600160601b63ffffffff60c01b03198387099550600160601b63ffffffff60c01b031986600160601b63ffffffff60c01b03190382089b50600160601b63ffffffff60c01b03198a84099a5050505050505050505050955095509592505050565b60405180606001604052806003906020820280388339509192915050565b6040518060a001604052806005906020820280388339509192915050565b60405180608001604052806004906020820280388339509192915050565b6040518060400160405280600290602082028038833950919291505056fea265627a7a72305820fee9f515c64b095280681f89fe4f105a5df80c2e4fc9c21289038be47752dbd864736f6c63430005090032`

// DeployMulSigP256 deploys a new Ethereum contract, binding an instance of MulSigP256 to it.
func DeployMulSigP256(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MulSigP256, error) {
	parsed, err := abi.JSON(strings.NewReader(MulSigP256ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MulSigP256Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MulSigP256{MulSigP256Caller: MulSigP256Caller{contract: contract}, MulSigP256Transactor: MulSigP256Transactor{contract: contract}, MulSigP256Filterer: MulSigP256Filterer{contract: contract}}, nil
}

// MulSigP256 is an auto generated Go binding around an Ethereum contract.
type MulSigP256 struct {
	MulSigP256Caller     // Read-only binding to the contract
	MulSigP256Transactor // Write-only binding to the contract
	MulSigP256Filterer   // Log filterer for contract events
}

// MulSigP256Caller is an auto generated read-only Go binding around an Ethereum contract.
type MulSigP256Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulSigP256Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MulSigP256Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulSigP256Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulSigP256Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulSigP256Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulSigP256Session struct {
	Contract     *MulSigP256       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulSigP256CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulSigP256CallerSession struct {
	Contract *MulSigP256Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MulSigP256TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulSigP256TransactorSession struct {
	Contract     *MulSigP256Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MulSigP256Raw is an auto generated low-level Go binding around an Ethereum contract.
type MulSigP256Raw struct {
	Contract *MulSigP256 // Generic contract binding to access the raw methods on
}

// MulSigP256CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulSigP256CallerRaw struct {
	Contract *MulSigP256Caller // Generic read-only contract binding to access the raw methods on
}

// MulSigP256TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulSigP256TransactorRaw struct {
	Contract *MulSigP256Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMulSigP256 creates a new instance of MulSigP256, bound to a specific deployed contract.
func NewMulSigP256(address common.Address, backend bind.ContractBackend) (*MulSigP256, error) {
	contract, err := bindMulSigP256(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MulSigP256{MulSigP256Caller: MulSigP256Caller{contract: contract}, MulSigP256Transactor: MulSigP256Transactor{contract: contract}, MulSigP256Filterer: MulSigP256Filterer{contract: contract}}, nil
}

// NewMulSigP256Caller creates a new read-only instance of MulSigP256, bound to a specific deployed contract.
func NewMulSigP256Caller(address common.Address, caller bind.ContractCaller) (*MulSigP256Caller, error) {
	contract, err := bindMulSigP256(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulSigP256Caller{contract: contract}, nil
}

// NewMulSigP256Transactor creates a new write-only instance of MulSigP256, bound to a specific deployed contract.
func NewMulSigP256Transactor(address common.Address, transactor bind.ContractTransactor) (*MulSigP256Transactor, error) {
	contract, err := bindMulSigP256(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulSigP256Transactor{contract: contract}, nil
}

// NewMulSigP256Filterer creates a new log filterer instance of MulSigP256, bound to a specific deployed contract.
func NewMulSigP256Filterer(address common.Address, filterer bind.ContractFilterer) (*MulSigP256Filterer, error) {
	contract, err := bindMulSigP256(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulSigP256Filterer{contract: contract}, nil
}

// bindMulSigP256 binds a generic wrapper to an already deployed contract.
func bindMulSigP256(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MulSigP256ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulSigP256 *MulSigP256Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MulSigP256.Contract.MulSigP256Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulSigP256 *MulSigP256Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulSigP256.Contract.MulSigP256Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulSigP256 *MulSigP256Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulSigP256.Contract.MulSigP256Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulSigP256 *MulSigP256CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MulSigP256.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulSigP256 *MulSigP256TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulSigP256.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulSigP256 *MulSigP256TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulSigP256.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0xe022d77c.
//
// Solidity: function add(uint256 x0, uint256 y0, uint256 x1, uint256 y1) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256Caller) Add(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int, x1 *big.Int, y1 *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MulSigP256.contract.Call(opts, out, "add", x0, y0, x1, y1)
	return *ret0, *ret1, err
}

// Add is a free data retrieval call binding the contract method 0xe022d77c.
//
// Solidity: function add(uint256 x0, uint256 y0, uint256 x1, uint256 y1) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256Session) Add(x0 *big.Int, y0 *big.Int, x1 *big.Int, y1 *big.Int) (*big.Int, *big.Int, error) {
	return _MulSigP256.Contract.Add(&_MulSigP256.CallOpts, x0, y0, x1, y1)
}

// Add is a free data retrieval call binding the contract method 0xe022d77c.
//
// Solidity: function add(uint256 x0, uint256 y0, uint256 x1, uint256 y1) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256CallerSession) Add(x0 *big.Int, y0 *big.Int, x1 *big.Int, y1 *big.Int) (*big.Int, *big.Int, error) {
	return _MulSigP256.Contract.Add(&_MulSigP256.CallOpts, x0, y0, x1, y1)
}

// AddAndReturnProjectivePoint is a free data retrieval call binding the contract method 0x713eca28.
//
// Solidity: function addAndReturnProjectivePoint(uint256 x1, uint256 y1, uint256 x2, uint256 y2) constant returns(uint256[3] P)
func (_MulSigP256 *MulSigP256Caller) AddAndReturnProjectivePoint(opts *bind.CallOpts, x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) ([3]*big.Int, error) {
	var (
		ret0 = new([3]*big.Int)
	)
	out := ret0
	err := _MulSigP256.contract.Call(opts, out, "addAndReturnProjectivePoint", x1, y1, x2, y2)
	return *ret0, err
}

// AddAndReturnProjectivePoint is a free data retrieval call binding the contract method 0x713eca28.
//
// Solidity: function addAndReturnProjectivePoint(uint256 x1, uint256 y1, uint256 x2, uint256 y2) constant returns(uint256[3] P)
func (_MulSigP256 *MulSigP256Session) AddAndReturnProjectivePoint(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) ([3]*big.Int, error) {
	return _MulSigP256.Contract.AddAndReturnProjectivePoint(&_MulSigP256.CallOpts, x1, y1, x2, y2)
}

// AddAndReturnProjectivePoint is a free data retrieval call binding the contract method 0x713eca28.
//
// Solidity: function addAndReturnProjectivePoint(uint256 x1, uint256 y1, uint256 x2, uint256 y2) constant returns(uint256[3] P)
func (_MulSigP256 *MulSigP256CallerSession) AddAndReturnProjectivePoint(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) ([3]*big.Int, error) {
	return _MulSigP256.Contract.AddAndReturnProjectivePoint(&_MulSigP256.CallOpts, x1, y1, x2, y2)
}

// AddProj is a free data retrieval call binding the contract method 0xf214aba0.
//
// Solidity: function addProj(uint256 x0, uint256 y0, uint256 z0, uint256 x1, uint256 y1, uint256 z1) constant returns(uint256 x2, uint256 y2, uint256 z2)
func (_MulSigP256 *MulSigP256Caller) AddProj(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int, z0 *big.Int, x1 *big.Int, y1 *big.Int, z1 *big.Int) (struct {
	X2 *big.Int
	Y2 *big.Int
	Z2 *big.Int
}, error) {
	ret := new(struct {
		X2 *big.Int
		Y2 *big.Int
		Z2 *big.Int
	})
	out := ret
	err := _MulSigP256.contract.Call(opts, out, "addProj", x0, y0, z0, x1, y1, z1)
	return *ret, err
}

// AddProj is a free data retrieval call binding the contract method 0xf214aba0.
//
// Solidity: function addProj(uint256 x0, uint256 y0, uint256 z0, uint256 x1, uint256 y1, uint256 z1) constant returns(uint256 x2, uint256 y2, uint256 z2)
func (_MulSigP256 *MulSigP256Session) AddProj(x0 *big.Int, y0 *big.Int, z0 *big.Int, x1 *big.Int, y1 *big.Int, z1 *big.Int) (struct {
	X2 *big.Int
	Y2 *big.Int
	Z2 *big.Int
}, error) {
	return _MulSigP256.Contract.AddProj(&_MulSigP256.CallOpts, x0, y0, z0, x1, y1, z1)
}

// AddProj is a free data retrieval call binding the contract method 0xf214aba0.
//
// Solidity: function addProj(uint256 x0, uint256 y0, uint256 z0, uint256 x1, uint256 y1, uint256 z1) constant returns(uint256 x2, uint256 y2, uint256 z2)
func (_MulSigP256 *MulSigP256CallerSession) AddProj(x0 *big.Int, y0 *big.Int, z0 *big.Int, x1 *big.Int, y1 *big.Int, z1 *big.Int) (struct {
	X2 *big.Int
	Y2 *big.Int
	Z2 *big.Int
}, error) {
	return _MulSigP256.Contract.AddProj(&_MulSigP256.CallOpts, x0, y0, z0, x1, y1, z1)
}

// CheckMulSig is a free data retrieval call binding the contract method 0x883c48d4.
//
// Solidity: function checkMulSig(uint256[8] xPks, uint256[8] yPks, uint256[8] idxSigs, int128 sigsLen, uint256 xR, uint256 yR, bytes bytesR, uint256 sig, bytes32 mess) constant returns(bool)
func (_MulSigP256 *MulSigP256Caller) CheckMulSig(opts *bind.CallOpts, xPks [8]*big.Int, yPks [8]*big.Int, idxSigs [8]*big.Int, sigsLen *big.Int, xR *big.Int, yR *big.Int, bytesR []byte, sig *big.Int, mess [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MulSigP256.contract.Call(opts, out, "checkMulSig", xPks, yPks, idxSigs, sigsLen, xR, yR, bytesR, sig, mess)
	return *ret0, err
}

// CheckMulSig is a free data retrieval call binding the contract method 0x883c48d4.
//
// Solidity: function checkMulSig(uint256[8] xPks, uint256[8] yPks, uint256[8] idxSigs, int128 sigsLen, uint256 xR, uint256 yR, bytes bytesR, uint256 sig, bytes32 mess) constant returns(bool)
func (_MulSigP256 *MulSigP256Session) CheckMulSig(xPks [8]*big.Int, yPks [8]*big.Int, idxSigs [8]*big.Int, sigsLen *big.Int, xR *big.Int, yR *big.Int, bytesR []byte, sig *big.Int, mess [32]byte) (bool, error) {
	return _MulSigP256.Contract.CheckMulSig(&_MulSigP256.CallOpts, xPks, yPks, idxSigs, sigsLen, xR, yR, bytesR, sig, mess)
}

// CheckMulSig is a free data retrieval call binding the contract method 0x883c48d4.
//
// Solidity: function checkMulSig(uint256[8] xPks, uint256[8] yPks, uint256[8] idxSigs, int128 sigsLen, uint256 xR, uint256 yR, bytes bytesR, uint256 sig, bytes32 mess) constant returns(bool)
func (_MulSigP256 *MulSigP256CallerSession) CheckMulSig(xPks [8]*big.Int, yPks [8]*big.Int, idxSigs [8]*big.Int, sigsLen *big.Int, xR *big.Int, yR *big.Int, bytesR []byte, sig *big.Int, mess [32]byte) (bool, error) {
	return _MulSigP256.Contract.CheckMulSig(&_MulSigP256.CallOpts, xPks, yPks, idxSigs, sigsLen, xR, yR, bytesR, sig, mess)
}

// GenCommonParams is a free data retrieval call binding the contract method 0x80a694de.
//
// Solidity: function genCommonParams(uint256[8] xPks, uint256[8] yPks, uint256[8] idxSigs, uint8 sigLen, bytes bytesR, bytes32 mess) constant returns(uint256[3] res)
func (_MulSigP256 *MulSigP256Caller) GenCommonParams(opts *bind.CallOpts, xPks [8]*big.Int, yPks [8]*big.Int, idxSigs [8]*big.Int, sigLen uint8, bytesR []byte, mess [32]byte) ([3]*big.Int, error) {
	var (
		ret0 = new([3]*big.Int)
	)
	out := ret0
	err := _MulSigP256.contract.Call(opts, out, "genCommonParams", xPks, yPks, idxSigs, sigLen, bytesR, mess)
	return *ret0, err
}

// GenCommonParams is a free data retrieval call binding the contract method 0x80a694de.
//
// Solidity: function genCommonParams(uint256[8] xPks, uint256[8] yPks, uint256[8] idxSigs, uint8 sigLen, bytes bytesR, bytes32 mess) constant returns(uint256[3] res)
func (_MulSigP256 *MulSigP256Session) GenCommonParams(xPks [8]*big.Int, yPks [8]*big.Int, idxSigs [8]*big.Int, sigLen uint8, bytesR []byte, mess [32]byte) ([3]*big.Int, error) {
	return _MulSigP256.Contract.GenCommonParams(&_MulSigP256.CallOpts, xPks, yPks, idxSigs, sigLen, bytesR, mess)
}

// GenCommonParams is a free data retrieval call binding the contract method 0x80a694de.
//
// Solidity: function genCommonParams(uint256[8] xPks, uint256[8] yPks, uint256[8] idxSigs, uint8 sigLen, bytes bytesR, bytes32 mess) constant returns(uint256[3] res)
func (_MulSigP256 *MulSigP256CallerSession) GenCommonParams(xPks [8]*big.Int, yPks [8]*big.Int, idxSigs [8]*big.Int, sigLen uint8, bytesR []byte, mess [32]byte) ([3]*big.Int, error) {
	return _MulSigP256.Contract.GenCommonParams(&_MulSigP256.CallOpts, xPks, yPks, idxSigs, sigLen, bytesR, mess)
}

// InverseMod is a free data retrieval call binding the contract method 0xca465e6e.
//
// Solidity: function inverseMod(uint256 num, uint256 modular) constant returns(uint256 k)
func (_MulSigP256 *MulSigP256Caller) InverseMod(opts *bind.CallOpts, num *big.Int, modular *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MulSigP256.contract.Call(opts, out, "inverseMod", num, modular)
	return *ret0, err
}

// InverseMod is a free data retrieval call binding the contract method 0xca465e6e.
//
// Solidity: function inverseMod(uint256 num, uint256 modular) constant returns(uint256 k)
func (_MulSigP256 *MulSigP256Session) InverseMod(num *big.Int, modular *big.Int) (*big.Int, error) {
	return _MulSigP256.Contract.InverseMod(&_MulSigP256.CallOpts, num, modular)
}

// InverseMod is a free data retrieval call binding the contract method 0xca465e6e.
//
// Solidity: function inverseMod(uint256 num, uint256 modular) constant returns(uint256 k)
func (_MulSigP256 *MulSigP256CallerSession) InverseMod(num *big.Int, modular *big.Int) (*big.Int, error) {
	return _MulSigP256.Contract.InverseMod(&_MulSigP256.CallOpts, num, modular)
}

// IsOnCurve is a free data retrieval call binding the contract method 0x0b0dbcfa.
//
// Solidity: function isOnCurve(uint256 x, uint256 y) constant returns(bool)
func (_MulSigP256 *MulSigP256Caller) IsOnCurve(opts *bind.CallOpts, x *big.Int, y *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MulSigP256.contract.Call(opts, out, "isOnCurve", x, y)
	return *ret0, err
}

// IsOnCurve is a free data retrieval call binding the contract method 0x0b0dbcfa.
//
// Solidity: function isOnCurve(uint256 x, uint256 y) constant returns(bool)
func (_MulSigP256 *MulSigP256Session) IsOnCurve(x *big.Int, y *big.Int) (bool, error) {
	return _MulSigP256.Contract.IsOnCurve(&_MulSigP256.CallOpts, x, y)
}

// IsOnCurve is a free data retrieval call binding the contract method 0x0b0dbcfa.
//
// Solidity: function isOnCurve(uint256 x, uint256 y) constant returns(bool)
func (_MulSigP256 *MulSigP256CallerSession) IsOnCurve(x *big.Int, y *big.Int) (bool, error) {
	return _MulSigP256.Contract.IsOnCurve(&_MulSigP256.CallOpts, x, y)
}

// IsZeroCurve is a free data retrieval call binding the contract method 0xc30cfa2d.
//
// Solidity: function isZeroCurve(uint256 x0, uint256 y0) constant returns(bool isZero)
func (_MulSigP256 *MulSigP256Caller) IsZeroCurve(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MulSigP256.contract.Call(opts, out, "isZeroCurve", x0, y0)
	return *ret0, err
}

// IsZeroCurve is a free data retrieval call binding the contract method 0xc30cfa2d.
//
// Solidity: function isZeroCurve(uint256 x0, uint256 y0) constant returns(bool isZero)
func (_MulSigP256 *MulSigP256Session) IsZeroCurve(x0 *big.Int, y0 *big.Int) (bool, error) {
	return _MulSigP256.Contract.IsZeroCurve(&_MulSigP256.CallOpts, x0, y0)
}

// IsZeroCurve is a free data retrieval call binding the contract method 0xc30cfa2d.
//
// Solidity: function isZeroCurve(uint256 x0, uint256 y0) constant returns(bool isZero)
func (_MulSigP256 *MulSigP256CallerSession) IsZeroCurve(x0 *big.Int, y0 *big.Int) (bool, error) {
	return _MulSigP256.Contract.IsZeroCurve(&_MulSigP256.CallOpts, x0, y0)
}

// MultipleGeneratorByScalar is a free data retrieval call binding the contract method 0x0afb4ddc.
//
// Solidity: function multipleGeneratorByScalar(uint256 scalar) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256Caller) MultipleGeneratorByScalar(opts *bind.CallOpts, scalar *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MulSigP256.contract.Call(opts, out, "multipleGeneratorByScalar", scalar)
	return *ret0, *ret1, err
}

// MultipleGeneratorByScalar is a free data retrieval call binding the contract method 0x0afb4ddc.
//
// Solidity: function multipleGeneratorByScalar(uint256 scalar) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256Session) MultipleGeneratorByScalar(scalar *big.Int) (*big.Int, *big.Int, error) {
	return _MulSigP256.Contract.MultipleGeneratorByScalar(&_MulSigP256.CallOpts, scalar)
}

// MultipleGeneratorByScalar is a free data retrieval call binding the contract method 0x0afb4ddc.
//
// Solidity: function multipleGeneratorByScalar(uint256 scalar) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256CallerSession) MultipleGeneratorByScalar(scalar *big.Int) (*big.Int, *big.Int, error) {
	return _MulSigP256.Contract.MultipleGeneratorByScalar(&_MulSigP256.CallOpts, scalar)
}

// MultiplyPowerBase2 is a free data retrieval call binding the contract method 0x322b24aa.
//
// Solidity: function multiplyPowerBase2(uint256 x0, uint256 y0, uint256 exp) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256Caller) MultiplyPowerBase2(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int, exp *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MulSigP256.contract.Call(opts, out, "multiplyPowerBase2", x0, y0, exp)
	return *ret0, *ret1, err
}

// MultiplyPowerBase2 is a free data retrieval call binding the contract method 0x322b24aa.
//
// Solidity: function multiplyPowerBase2(uint256 x0, uint256 y0, uint256 exp) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256Session) MultiplyPowerBase2(x0 *big.Int, y0 *big.Int, exp *big.Int) (*big.Int, *big.Int, error) {
	return _MulSigP256.Contract.MultiplyPowerBase2(&_MulSigP256.CallOpts, x0, y0, exp)
}

// MultiplyPowerBase2 is a free data retrieval call binding the contract method 0x322b24aa.
//
// Solidity: function multiplyPowerBase2(uint256 x0, uint256 y0, uint256 exp) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256CallerSession) MultiplyPowerBase2(x0 *big.Int, y0 *big.Int, exp *big.Int) (*big.Int, *big.Int, error) {
	return _MulSigP256.Contract.MultiplyPowerBase2(&_MulSigP256.CallOpts, x0, y0, exp)
}

// MultiplyScalar is a free data retrieval call binding the contract method 0x72fb4a14.
//
// Solidity: function multiplyScalar(uint256 x, uint256 y, uint256 k) constant returns(uint256 x1, uint256 y1)
func (_MulSigP256 *MulSigP256Caller) MultiplyScalar(opts *bind.CallOpts, x *big.Int, y *big.Int, k *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	ret := new(struct {
		X1 *big.Int
		Y1 *big.Int
	})
	out := ret
	err := _MulSigP256.contract.Call(opts, out, "multiplyScalar", x, y, k)
	return *ret, err
}

// MultiplyScalar is a free data retrieval call binding the contract method 0x72fb4a14.
//
// Solidity: function multiplyScalar(uint256 x, uint256 y, uint256 k) constant returns(uint256 x1, uint256 y1)
func (_MulSigP256 *MulSigP256Session) MultiplyScalar(x *big.Int, y *big.Int, k *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	return _MulSigP256.Contract.MultiplyScalar(&_MulSigP256.CallOpts, x, y, k)
}

// MultiplyScalar is a free data retrieval call binding the contract method 0x72fb4a14.
//
// Solidity: function multiplyScalar(uint256 x, uint256 y, uint256 k) constant returns(uint256 x1, uint256 y1)
func (_MulSigP256 *MulSigP256CallerSession) MultiplyScalar(x *big.Int, y *big.Int, k *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	return _MulSigP256.Contract.MultiplyScalar(&_MulSigP256.CallOpts, x, y, k)
}

// ToAffinePoint is a free data retrieval call binding the contract method 0x675ca043.
//
// Solidity: function toAffinePoint(uint256 x0, uint256 y0, uint256 z0) constant returns(uint256 x1, uint256 y1)
func (_MulSigP256 *MulSigP256Caller) ToAffinePoint(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int, z0 *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	ret := new(struct {
		X1 *big.Int
		Y1 *big.Int
	})
	out := ret
	err := _MulSigP256.contract.Call(opts, out, "toAffinePoint", x0, y0, z0)
	return *ret, err
}

// ToAffinePoint is a free data retrieval call binding the contract method 0x675ca043.
//
// Solidity: function toAffinePoint(uint256 x0, uint256 y0, uint256 z0) constant returns(uint256 x1, uint256 y1)
func (_MulSigP256 *MulSigP256Session) ToAffinePoint(x0 *big.Int, y0 *big.Int, z0 *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	return _MulSigP256.Contract.ToAffinePoint(&_MulSigP256.CallOpts, x0, y0, z0)
}

// ToAffinePoint is a free data retrieval call binding the contract method 0x675ca043.
//
// Solidity: function toAffinePoint(uint256 x0, uint256 y0, uint256 z0) constant returns(uint256 x1, uint256 y1)
func (_MulSigP256 *MulSigP256CallerSession) ToAffinePoint(x0 *big.Int, y0 *big.Int, z0 *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
}, error) {
	return _MulSigP256.Contract.ToAffinePoint(&_MulSigP256.CallOpts, x0, y0, z0)
}

// ToProjectivePoint is a free data retrieval call binding the contract method 0x7ec8da8d.
//
// Solidity: function toProjectivePoint(uint256 x0, uint256 y0) constant returns(uint256[3] P)
func (_MulSigP256 *MulSigP256Caller) ToProjectivePoint(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int) ([3]*big.Int, error) {
	var (
		ret0 = new([3]*big.Int)
	)
	out := ret0
	err := _MulSigP256.contract.Call(opts, out, "toProjectivePoint", x0, y0)
	return *ret0, err
}

// ToProjectivePoint is a free data retrieval call binding the contract method 0x7ec8da8d.
//
// Solidity: function toProjectivePoint(uint256 x0, uint256 y0) constant returns(uint256[3] P)
func (_MulSigP256 *MulSigP256Session) ToProjectivePoint(x0 *big.Int, y0 *big.Int) ([3]*big.Int, error) {
	return _MulSigP256.Contract.ToProjectivePoint(&_MulSigP256.CallOpts, x0, y0)
}

// ToProjectivePoint is a free data retrieval call binding the contract method 0x7ec8da8d.
//
// Solidity: function toProjectivePoint(uint256 x0, uint256 y0) constant returns(uint256[3] P)
func (_MulSigP256 *MulSigP256CallerSession) ToProjectivePoint(x0 *big.Int, y0 *big.Int) ([3]*big.Int, error) {
	return _MulSigP256.Contract.ToProjectivePoint(&_MulSigP256.CallOpts, x0, y0)
}

// Twice is a free data retrieval call binding the contract method 0x09d3ef31.
//
// Solidity: function twice(uint256 x0, uint256 y0) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256Caller) Twice(opts *bind.CallOpts, x0 *big.Int, y0 *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MulSigP256.contract.Call(opts, out, "twice", x0, y0)
	return *ret0, *ret1, err
}

// Twice is a free data retrieval call binding the contract method 0x09d3ef31.
//
// Solidity: function twice(uint256 x0, uint256 y0) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256Session) Twice(x0 *big.Int, y0 *big.Int) (*big.Int, *big.Int, error) {
	return _MulSigP256.Contract.Twice(&_MulSigP256.CallOpts, x0, y0)
}

// Twice is a free data retrieval call binding the contract method 0x09d3ef31.
//
// Solidity: function twice(uint256 x0, uint256 y0) constant returns(uint256, uint256)
func (_MulSigP256 *MulSigP256CallerSession) Twice(x0 *big.Int, y0 *big.Int) (*big.Int, *big.Int, error) {
	return _MulSigP256.Contract.Twice(&_MulSigP256.CallOpts, x0, y0)
}

// TwiceProj is a free data retrieval call binding the contract method 0xc80edca4.
//
// Solidity: function twiceProj(uint256 x, uint256 y, uint256 z) constant returns(uint256 x1, uint256 y1, uint256 z1)
func (_MulSigP256 *MulSigP256Caller) TwiceProj(opts *bind.CallOpts, x *big.Int, y *big.Int, z *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
	Z1 *big.Int
}, error) {
	ret := new(struct {
		X1 *big.Int
		Y1 *big.Int
		Z1 *big.Int
	})
	out := ret
	err := _MulSigP256.contract.Call(opts, out, "twiceProj", x, y, z)
	return *ret, err
}

// TwiceProj is a free data retrieval call binding the contract method 0xc80edca4.
//
// Solidity: function twiceProj(uint256 x, uint256 y, uint256 z) constant returns(uint256 x1, uint256 y1, uint256 z1)
func (_MulSigP256 *MulSigP256Session) TwiceProj(x *big.Int, y *big.Int, z *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
	Z1 *big.Int
}, error) {
	return _MulSigP256.Contract.TwiceProj(&_MulSigP256.CallOpts, x, y, z)
}

// TwiceProj is a free data retrieval call binding the contract method 0xc80edca4.
//
// Solidity: function twiceProj(uint256 x, uint256 y, uint256 z) constant returns(uint256 x1, uint256 y1, uint256 z1)
func (_MulSigP256 *MulSigP256CallerSession) TwiceProj(x *big.Int, y *big.Int, z *big.Int) (struct {
	X1 *big.Int
	Y1 *big.Int
	Z1 *big.Int
}, error) {
	return _MulSigP256.Contract.TwiceProj(&_MulSigP256.CallOpts, x, y, z)
}

// ZeroAffine is a free data retrieval call binding the contract method 0x84dfba46.
//
// Solidity: function zeroAffine() constant returns(uint256 x, uint256 y)
func (_MulSigP256 *MulSigP256Caller) ZeroAffine(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _MulSigP256.contract.Call(opts, out, "zeroAffine")
	return *ret, err
}

// ZeroAffine is a free data retrieval call binding the contract method 0x84dfba46.
//
// Solidity: function zeroAffine() constant returns(uint256 x, uint256 y)
func (_MulSigP256 *MulSigP256Session) ZeroAffine() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _MulSigP256.Contract.ZeroAffine(&_MulSigP256.CallOpts)
}

// ZeroAffine is a free data retrieval call binding the contract method 0x84dfba46.
//
// Solidity: function zeroAffine() constant returns(uint256 x, uint256 y)
func (_MulSigP256 *MulSigP256CallerSession) ZeroAffine() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _MulSigP256.Contract.ZeroAffine(&_MulSigP256.CallOpts)
}

// ZeroProj is a free data retrieval call binding the contract method 0x14c67060.
//
// Solidity: function zeroProj() constant returns(uint256 x, uint256 y, uint256 z)
func (_MulSigP256 *MulSigP256Caller) ZeroProj(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
	Z *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
		Z *big.Int
	})
	out := ret
	err := _MulSigP256.contract.Call(opts, out, "zeroProj")
	return *ret, err
}

// ZeroProj is a free data retrieval call binding the contract method 0x14c67060.
//
// Solidity: function zeroProj() constant returns(uint256 x, uint256 y, uint256 z)
func (_MulSigP256 *MulSigP256Session) ZeroProj() (struct {
	X *big.Int
	Y *big.Int
	Z *big.Int
}, error) {
	return _MulSigP256.Contract.ZeroProj(&_MulSigP256.CallOpts)
}

// ZeroProj is a free data retrieval call binding the contract method 0x14c67060.
//
// Solidity: function zeroProj() constant returns(uint256 x, uint256 y, uint256 z)
func (_MulSigP256 *MulSigP256CallerSession) ZeroProj() (struct {
	X *big.Int
	Y *big.Int
	Z *big.Int
}, error) {
	return _MulSigP256.Contract.ZeroProj(&_MulSigP256.CallOpts)
}
