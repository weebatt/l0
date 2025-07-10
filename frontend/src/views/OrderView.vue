<template>
  <v-container class="py-8">
    <v-card class="mx-auto" max-width="600">
      <v-card-title>Order Lookup</v-card-title>
      <v-card-text>
        <v-form @submit.prevent="fetchOrder">
          <v-text-field
            v-model="orderId"
            label="Order ID"
            required
            prepend-inner-icon="mdi-magnify"
          />
          <v-btn type="submit" color="primary" :loading="loading" class="mt-2">
            Get Order
          </v-btn>
        </v-form>
        <v-alert v-if="error" type="error" class="mt-4">{{ error }}</v-alert>
        <v-alert v-if="order" type="success" class="mt-4">
          <div><strong>Order UID:</strong> {{ order.order_uid }}</div>
          <div><strong>Track Number:</strong> {{ order.track_number }}</div>
          <div><strong>Entry:</strong> {{ order.entry }}</div>
          <div><strong>Locale:</strong> {{ order.locale }}</div>
          <div><strong>Customer ID:</strong> {{ order.customer_id }}</div>
          <div><strong>Delivery Service:</strong> {{ order.delivery_service }}</div>
          <div><strong>Shardkey:</strong> {{ order.shardkey }}</div>
          <div><strong>SM ID:</strong> {{ order.sm_id }}</div>
          <div><strong>Date Created:</strong> {{ order.date_created }}</div>
          <div><strong>Oof Shard:</strong> {{ order.oof_shard }}</div>
          <div><strong>Internal Signature:</strong> {{ order.internal_signature }}</div>
          <v-divider class="my-2" />
          <div>
            <strong>Delivery:</strong>
            <ul>
              <li><strong>Name:</strong> {{ order.delivery?.name }}</li>
              <li><strong>Phone:</strong> {{ order.delivery?.phone }}</li>
              <li><strong>Zip:</strong> {{ order.delivery?.zip }}</li>
              <li><strong>City:</strong> {{ order.delivery?.city }}</li>
              <li><strong>Address:</strong> {{ order.delivery?.address }}</li>
              <li><strong>Region:</strong> {{ order.delivery?.region }}</li>
              <li><strong>Email:</strong> {{ order.delivery?.email }}</li>
            </ul>
          </div>
          <v-divider class="my-2" />
          <div>
            <strong>Payment:</strong>
            <ul>
              <li><strong>Transaction:</strong> {{ order.payment?.transaction }}</li>
              <li><strong>Request ID:</strong> {{ order.payment?.request_id }}</li>
              <li><strong>Currency:</strong> {{ order.payment?.currency }}</li>
              <li><strong>Provider:</strong> {{ order.payment?.provider }}</li>
              <li><strong>Amount:</strong> {{ order.payment?.amount }}</li>
              <li><strong>Payment DT:</strong> {{ order.payment?.payment_dt }}</li>
              <li><strong>Bank:</strong> {{ order.payment?.bank }}</li>
              <li><strong>Delivery Cost:</strong> {{ order.payment?.delivery_cost }}</li>
              <li><strong>Goods Total:</strong> {{ order.payment?.goods_total }}</li>
              <li><strong>Custom Fee:</strong> {{ order.payment?.custom_fee }}</li>
            </ul>
          </div>
          <v-divider class="my-2" />
          <div>
            <strong>Items:</strong>
            <v-list v-if="order.items && order.items.length">
              <v-list-item v-for="(item, idx) in order.items" :key="idx">
                <div>
                  <strong>chrt_id:</strong> {{ item.chrt_id }},
                  <strong>track_number:</strong> {{ item.track_number }},
                  <strong>price:</strong> {{ item.price }},
                  <strong>rid:</strong> {{ item.rid }},
                  <strong>name:</strong> {{ item.name }},
                  <strong>sale:</strong> {{ item.sale }},
                  <strong>size:</strong> {{ item.size }},
                  <strong>total_price:</strong> {{ item.total_price }},
                  <strong>nm_id:</strong> {{ item.nm_id }},
                  <strong>brand:</strong> {{ item.brand }},
                  <strong>status:</strong> {{ item.status }}
                </div>
              </v-list-item>
            </v-list>
            <div v-else>No items</div>
          </div>
        </v-alert>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'

const orderId = ref('')
const order = ref(null)
const loading = ref(false)
const error = ref('')
const route = useRoute()

async function fetchOrder() {
  loading.value = true
  error.value = ''
  order.value = null
  try {
    const res = await axios.get(`http://localhost:8081/order?order_uid=${orderId.value}`)
    order.value = res.data
  } catch (e) {
    error.value = 'Order not found or error fetching order.'
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  if (route.params.order_uid) {
    orderId.value = route.params.order_uid
    fetchOrder()
  }
})
</script>