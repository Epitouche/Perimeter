# Comparative Study Summary

## Table of Contents

- [Comparative Study Summary](#comparative-study-summary)
  - [Table of Contents](#table-of-contents)
  - [Frontend Technology](#frontend-technology)
    - [Angular vs Nuxt vs Remix](#angular-vs-nuxt-vs-remix)
  - [Database Systems](#database-systems)
    - [PostgreSQL vs MongoDB vs MySQL](#postgresql-vs-mongodb-vs-mysql)
  - [Mobile Development](#mobile-development)
    - [Kotlin vs React Native](#kotlin-vs-react-native)
  - [Backend Technology](#backend-technology)
    - [Golang vs Gleam vs Java](#golang-vs-gleam-vs-java)
  - [Overall Conclusion](#overall-conclusion)

---

## Frontend Technology

### Angular vs Nuxt vs Remix

| Criteria                | Angular                          | Nuxt (Vue.js)                   | Remix (React)                   |
|-------------------------|----------------------------------|----------------------------------|----------------------------------|
| **Learning Curve**      | Steep                            | Moderate                         | Moderate                         |
| **Performance**         | Good                             | Excellent (SSR)                 | Optimized (SSR)                 |
| **Accessibility**       | Built-in ARIA support with Angular CDK<br>Keyboard navigation | Vue A11y plugin for ARIA<br>Semantic HTML with SSR | React ARIA libraries<br>Custom `aria-*` attributes for screen readers |
| **Ecosystem**           | Mature                           | Growing                          | Emerging                        |
| **Best For**            | Enterprise apps                  | Medium projects                  | Modern web apps                 |

**Choice**: Nuxt  
**Why**:  

- Balanced accessibility with plugins like Vue A11y  
- Server-side rendering improves compatibility with assistive technologies  
- Easy integration of semantic HTML for screen readers  

---

## Database Systems

### PostgreSQL vs MongoDB vs MySQL

| Criteria                | PostgreSQL                      | MongoDB                        | MySQL                          |
|-------------------------|--------------------------------|--------------------------------|--------------------------------|
| **Data Model**          | Relational + JSONB             | Document-based                 | Relational                    |
| **Scalability**         | Vertical + extensions          | Horizontal                     | Vertical                      |
| **Accessibility**       | Stores accessibility metadata in JSONB<br>Supports full-text search for screen-reader content | Flexible schema for storing user preferences<br>Supports location-based features for disabilities | Structured schemas enforce consistent accessibility data<br>Limited native tools for assistive needs |

**Choice**: PostgreSQL  
**Why**:  

- JSONB enables storing metadata for accessibility (e.g., alt text, captions)  
- Full-text search supports queries for screen-reader-friendly content  
- Extensible for advanced accessibility use cases  

---

## Mobile Development

### Kotlin vs React Native

| Criteria                | Kotlin                          | React Native                   |
|-------------------------|--------------------------------|--------------------------------|
| **Performance**         | Native                         | Bridge-dependent              |
| **Code Sharing**        | KMM (logic only)               | Full cross-platform           |
| **Accessibility**       | Platform-specific APIs for TalkBack (Android) and VoiceOver (iOS) | Built-in props: `accessibilityLabel`, `accessibilityHint`, `accessible` |
| **Ecosystem**           | Android-focused                | Mature cross-platform         |

**Choice**: React Native  
**Why**:

- Native props (`accessibilityLabel`, `accessibilityHint`) simplify implementation for screen readers
- Cross-platform development ensures consistent accessibility across Android and iOS  
- Strong community support for accessible UI patterns  

---

## Backend Technology

### Golang vs Gleam vs Java

| Criteria                | Golang                         | Gleam                         | Java                          |
|-------------------------|-------------------------------|-------------------------------|-------------------------------|
| **Concurrency**         | Goroutines                    | Erlang/OTP                   | Threads                      |
| **Deployment**          | Static binaries               | Container-friendly           | JVM-based                    |
| **Learning Curve**      | Moderate                      | Steep (FP)                   | Steep                        |
| **Accessibility**       | Efficient real-time APIs for assistive tools<br>Native compilation avoids latency for accessibility services | Fault-tolerant systems<br>Customizable workflows for accessibility-critical features | Robust frameworks (e.g., Spring) support WCAG compliance and microservices for assistive needs |

**Choice**: Golang  
**Why**:

- Lightweight and fast APIs for assistive technology integration  
- Goroutines handle real-time data for accessibility services (e.g., live captions)  
- Cloud-native and efficient for deploying scalable accessible systems  

---

## Overall Conclusion

**AREA Project Tech Stack**:

- Frontend: Nuxt (Vue.js)  
- Database: PostgreSQL  
- Mobile: React Native  
- Backend: Golang  

**Accessibility Strategy**:

1. **Frontend**: Nuxt leverages semantic HTML and ARIA standards for screen reader compatibility.  
2. **Database**: PostgreSQL stores accessibility metadata (e.g., alt text, captions) in JSONB and supports advanced search for assistive tools.  
3. **Mobile**: React Native uses `accessibilityLabel`, `accessibilityHint`, and `accessible` props for TalkBack and VoiceOver support.  
4. **Backend**: Golang ensures low-latency APIs for real-time accessibility features like live captions and notifications.  

This stack prioritizes inclusivity and ease of use for users with disabilities while maintaining high performance and scalability.
