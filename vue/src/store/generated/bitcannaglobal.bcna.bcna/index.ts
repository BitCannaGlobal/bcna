import { Client, registry, MissingWalletError } from 'BitCannaGlobal-bcna-client-ts'

import { Bitcannaid } from "BitCannaGlobal-bcna-client-ts/bitcannaglobal.bcna.bcna/types"
import { Params } from "BitCannaGlobal-bcna-client-ts/bitcannaglobal.bcna.bcna/types"
import { Supplychain } from "BitCannaGlobal-bcna-client-ts/bitcannaglobal.bcna.bcna/types"


export { Bitcannaid, Params, Supplychain };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				Bitcannaid: {},
				BitcannaidAll: {},
				BitcannaidByBcnaid: {},
				Supplychain: {},
				SupplychainAll: {},
				
				_Structure: {
						Bitcannaid: getStructure(Bitcannaid.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						Supplychain: getStructure(Supplychain.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getBitcannaid: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Bitcannaid[JSON.stringify(params)] ?? {}
		},
				getBitcannaidAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.BitcannaidAll[JSON.stringify(params)] ?? {}
		},
				getBitcannaidByBcnaid: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.BitcannaidByBcnaid[JSON.stringify(params)] ?? {}
		},
				getSupplychain: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Supplychain[JSON.stringify(params)] ?? {}
		},
				getSupplychainAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SupplychainAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: bitcannaglobal.bcna.bcna initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.BitcannaglobalBcnaBcna.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBitcannaid({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.BitcannaglobalBcnaBcna.query.queryBitcannaid( key.id)).data
				
					
				commit('QUERY', { query: 'Bitcannaid', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBitcannaid', payload: { options: { all }, params: {...key},query }})
				return getters['getBitcannaid']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBitcannaid API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBitcannaidAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.BitcannaglobalBcnaBcna.query.queryBitcannaidAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.BitcannaglobalBcnaBcna.query.queryBitcannaidAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'BitcannaidAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBitcannaidAll', payload: { options: { all }, params: {...key},query }})
				return getters['getBitcannaidAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBitcannaidAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBitcannaidByBcnaid({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.BitcannaglobalBcnaBcna.query.queryBitcannaidByBcnaid( key.bcnaid)).data
				
					
				commit('QUERY', { query: 'BitcannaidByBcnaid', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBitcannaidByBcnaid', payload: { options: { all }, params: {...key},query }})
				return getters['getBitcannaidByBcnaid']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBitcannaidByBcnaid API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySupplychain({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.BitcannaglobalBcnaBcna.query.querySupplychain( key.id)).data
				
					
				commit('QUERY', { query: 'Supplychain', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySupplychain', payload: { options: { all }, params: {...key},query }})
				return getters['getSupplychain']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySupplychain API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySupplychainAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.BitcannaglobalBcnaBcna.query.querySupplychainAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.BitcannaglobalBcnaBcna.query.querySupplychainAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'SupplychainAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySupplychainAll', payload: { options: { all }, params: {...key},query }})
				return getters['getSupplychainAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySupplychainAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateBitcannaid({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.BitcannaglobalBcnaBcna.tx.sendMsgCreateBitcannaid({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateBitcannaid:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateBitcannaid:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateBitcannaid({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.BitcannaglobalBcnaBcna.tx.sendMsgUpdateBitcannaid({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateBitcannaid:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateBitcannaid:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateSupplychain({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.BitcannaglobalBcnaBcna.tx.sendMsgUpdateSupplychain({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateSupplychain:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateSupplychain:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteBitcannaid({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.BitcannaglobalBcnaBcna.tx.sendMsgDeleteBitcannaid({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteBitcannaid:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteBitcannaid:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateSupplychain({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.BitcannaglobalBcnaBcna.tx.sendMsgCreateSupplychain({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateSupplychain:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateSupplychain:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteSupplychain({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.BitcannaglobalBcnaBcna.tx.sendMsgDeleteSupplychain({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteSupplychain:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteSupplychain:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateBitcannaid({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.BitcannaglobalBcnaBcna.tx.msgCreateBitcannaid({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateBitcannaid:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateBitcannaid:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateBitcannaid({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.BitcannaglobalBcnaBcna.tx.msgUpdateBitcannaid({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateBitcannaid:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateBitcannaid:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateSupplychain({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.BitcannaglobalBcnaBcna.tx.msgUpdateSupplychain({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateSupplychain:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateSupplychain:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteBitcannaid({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.BitcannaglobalBcnaBcna.tx.msgDeleteBitcannaid({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteBitcannaid:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteBitcannaid:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateSupplychain({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.BitcannaglobalBcnaBcna.tx.msgCreateSupplychain({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateSupplychain:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateSupplychain:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteSupplychain({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.BitcannaglobalBcnaBcna.tx.msgDeleteSupplychain({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteSupplychain:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteSupplychain:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}